package jitdec

import (
	"encoding/json"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/bytedance/sonic/internal/decoder/consts"
	"github.com/bytedance/sonic/internal/decoder/errors"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/loader"
	"github.com/bytedance/sonic/option"
	"github.com/bytedance/sonic/utf8"
)

type (
	MismatchTypeError = errors.MismatchTypeError
	SyntaxError       = errors.SyntaxError
)

const (
	_F_allow_control    = consts.F_allow_control
	_F_copy_string      = consts.F_copy_string
	_F_disable_unknown  = consts.F_disable_unknown
	_F_disable_urc      = consts.F_disable_urc
	_F_use_int64        = consts.F_use_int64
	_F_use_number       = consts.F_use_number
	_F_no_validate_json = consts.F_no_validate_json
	_F_validate_string  = consts.F_validate_string
	_F_case_sensitive   = consts.F_case_sensitive
)

var (
	error_wrap     = errors.ErrorWrap
	error_type     = errors.ErrorType
	error_field    = errors.ErrorField
	error_value    = errors.ErrorValue
	error_mismatch = errors.ErrorMismatch
	stackOverflow  = errors.StackOverflow
)

var decoderJitLoader = loader.Loader{
	Name: "sonic.jit.",
	File: "github.com/bytedance/sonic/jit.go",
	Options: loader.Options{
		NoPreempt: true,
	},
}

// Decode parses the JSON-encoded data from current position and stores the result
// in the value pointed to by val.
func Decode(s *string, i *int, f uint64, val interface{}) error {
	/* validate json if needed */
	if (f&(1<<_F_validate_string)) != 0 && !utf8.ValidateString(*s) {
		dbuf := utf8.CorrectWith(nil, rt.Str2Mem(*s), "\ufffd")
		*s = rt.Mem2Str(dbuf)
	}

	vv := rt.UnpackEface(val)
	vp := vv.Value

	/* check for nil type */
	if vv.Type == nil {
		return &json.InvalidUnmarshalError{}
	}

	/* must be a non-nil pointer */
	if vp == nil || vv.Type.Kind() != reflect.Ptr {
		return &json.InvalidUnmarshalError{Type: vv.Type.Pack()}
	}

	etp := rt.PtrElem(vv.Type)

	/* check the defined pointer type for issue 379 */
	if vv.Type.IsNamed() {
		newp := vp
		etp = vv.Type
		vp = unsafe.Pointer(&newp)
	}

	/* create a new stack, and call the decoder */
	sb := newStack()
	nb, err := decodeTypedPointer(*s, *i, etp, vp, sb, f)
	/* return the stack back */
	*i = nb
	freeStack(sb)

	/* avoid GC ahead */
	runtime.KeepAlive(vv)
	return err
}

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
	return PretouchMany([]reflect.Type{vt}, opts...)
}

// PretouchMany compiles all vts ahead-of-time to avoid JIT compilation on-the-fly,
// in order to reduce the first-hit latency.
func PretouchMany(vts []reflect.Type, opts ...option.CompileOption) error {
	if len(vts) == 0 {
		return nil
	}

	cfg := option.DefaultCompileOptions()
	for _, opt := range opts {
		opt(&cfg)
	}

	vtm := make(map[reflect.Type]bool, len(vts))
	for _, vt := range vts {
		vtm[vt] = true
	}
	return pretouchRec(vtm, cfg)
}

type jitdecPretouchProgram struct {
	vt   *rt.GoType
	item loader.LoadOneItem
}

func pretouchType(_vt reflect.Type, opts option.CompileOptions) (map[reflect.Type]bool, error) {
	/* compile function */
	compiler := newCompiler().apply(opts)
	decoder := func(vt *rt.GoType, _ ...interface{}) (interface{}, error) {
		if pp, err := compiler.compile(_vt); err != nil {
			return nil, err
		} else {
			as := newAssembler(pp)
			as.name = _vt.String()
			return as.Load(), nil
		}
	}

	/* find or compile */
	vt := rt.UnpackType(_vt)
	if val := programCache.Get(vt); val != nil {
		return nil, nil
	} else if _, err := programCache.Compute(vt, decoder); err == nil {
		return compiler.rec, nil
	} else {
		return nil, err
	}
}

func pretouchRec(vtm map[reflect.Type]bool, opts option.CompileOptions) error {
	pendings := make(map[*rt.GoType]jitdecPretouchProgram)

	for opts.RecursiveDepth >= 0 && len(vtm) > 0 {
		next := make(map[reflect.Type]bool)
		for vt := range vtm {
			gvt := rt.UnpackType(vt)
			if programCache.Get(gvt) != nil {
				continue
			}
			if _, ok := pendings[gvt]; ok {
				continue
			}

			compiler := newCompiler().apply(opts)
			pp, err := compiler.compile(vt)
			if err != nil {
				return err
			}

			as := newAssembler(pp)
			as.name = vt.String()
			text, pcdata := as.BaseAssembler.Export()

			pendings[gvt] = jitdecPretouchProgram{
				vt: gvt,
				item: loader.LoadOneItem{
					Text:      text,
					FuncName:  "decode_" + as.name,
					ArgSize:   _FP_args,
					ArgPtrs:   argPtrs,
					LocalPtrs: localPtrs,
					Pcdata:    pcdata,
				},
			}

			for svt := range compiler.rec {
				next[svt] = true
			}
		}

		opts.RecursiveDepth--
		vtm = next
	}

	if len(pendings) == 0 {
		return nil
	}

	entries := make([]jitdecPretouchProgram, 0, len(pendings))
	items := make([]loader.LoadOneItem, 0, len(pendings))
	for _, p := range pendings {
		entries = append(entries, p)
		items = append(items, p.item)
	}

	loaded := decoderJitLoader.LoadMany(items)
	for i, p := range entries {
		dec := ptodec(loaded[i])
		_, err := programCache.Compute(p.vt, func(*rt.GoType, ...interface{}) (interface{}, error) {
			return dec, nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
