package decoder

import (
	"reflect"
	"unsafe"

	"encoding/json"
	"github.com/bytedance/sonic/ast"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/internal"
	"github.com/bytedance/sonic/option"
)

// Decoder is the decoder context object
type Decoder struct {
	json  string
	opts  Options
	pos  int
}

// NewDecoder creates a new decoder instance.
func NewDecoder(s string) *Decoder {
	return &Decoder{json: s}
}

// Decode parses the JSON-encoded data from current position and stores the result
// in the value pointed to by val.
func (self *Decoder) Decode(val interface{}) error {
	return self.decodeImpl(val)
}

func (self *Decoder) decodeImpl(val interface{}) error {
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

	dec, err := findOrCompile(etp)
	if err != nil {
		return err
	}

	/* parse into document */
	ctx, err := internal.NewContext(self.json, self.pos, uint64(self.opts), etp)
	defer ctx.Delete()
	if ctx.Parser.Utf8Inv {
		self.json = ctx.Parser.Json
	}
	if err != nil {
		goto fix_error;
	}
	err = dec.FromDom(vp, ctx.Root(), &ctx)

fix_error:
	err = self.fix_error(err)

	// update position at last
	self.pos += ctx.Parser.Pos()
	return err
}

func (self *Decoder) fix_error(err error) error {

	if e, ok := err.(*internal.ParseError); ok {
		return SyntaxError{
			Pos: int(e.Offset) + self.pos,
			Src: self.json,
			Msg: e.Msg,
		}
	}

	if e, ok := err.(*internal.UnmatchedError); ok {
		return &MismatchTypeError {
			Pos: int(e.Offset) + self.pos,
			Src: self.json,
			Type: e.Type.Pack(),
		}
	}

	return err
}

func (self *Decoder) Reset(s string) {
    self.json = s
	self.pos = 0
}

func (self *Decoder) Pos() int {
	return self.pos
}

func (self *Decoder) CheckTrailings() error {
	for self.pos < len(self.json) && isSpace(self.json[self.pos]) {
		self.pos ++;
	}
	if self.pos < len(self.json) {
		return SyntaxError{
			Pos: self.pos,
			Src: self.json,
			Msg: "trailing characters",
		}
	}
	return nil
}

// Skip skips only one json value, and returns first non-blank character position and its ending position if it is valid.
// Otherwise, returns negative error code using start and invalid character position using end
func Skip(data []byte) (start int, end int) {
	pos := 0
	start, err := ast.Skip(rt.Mem2Str(data), &pos)
	if err != nil {
		return -1, pos
	}
	return start, pos
}

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
    cfg := option.DefaultCompileOptions()
    for _, opt := range opts {
        opt(&cfg)
    }
    return pretouchRec(map[reflect.Type]bool{vt:true}, cfg)
}

func pretouchType(_vt reflect.Type, opts option.CompileOptions) (map[reflect.Type]bool, error) {
    /* compile function */
    compiler := newCompiler().apply(opts)
    decoder := func(vt *rt.GoType, _ ...interface{}) (interface{}, error) {
        if f, err := compiler.compileType(_vt); err != nil {
            return nil, err
        } else {
            return f, nil
        }
    }

    /* find or compile */
    vt := rt.UnpackType(_vt)
    if val := programCache.Get(vt); val != nil {
        return nil, nil
    } else if _, err := programCache.Compute(vt, decoder); err == nil {
        return compiler.visited, nil
    } else {
        return nil, err
    }
}

func pretouchRec(vtm map[reflect.Type]bool, opts option.CompileOptions) error {
    if opts.RecursiveDepth < 0 || len(vtm) == 0 {
        return nil
    }
    next := make(map[reflect.Type]bool)
    for vt := range(vtm) {
        sub, err := pretouchType(vt, opts)
        if err != nil {
            return err
        }
        for svt := range(sub) {
            next[svt] = true
        }
    }
    opts.RecursiveDepth -= 1
    return pretouchRec(next, opts)
}
