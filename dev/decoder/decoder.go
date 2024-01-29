package decoder

import (
	"reflect"
	"unsafe"

	"encoding/json"
	"github.com/bytedance/sonic/dev/internal/rt"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

// Decoder is the decoder context object
type Decoder struct {
	json string
	opts Options
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
	/* parse into document */
	ctx, err := newCtx(self.json, self.opts)
	defer ctx.Context.Delete()
	if err != nil {
		return error_parse_internal(err, self.json)
	}

	/* check utf8 lossy and copy the repred string  */
	if ctx.Dom.HasUtf8Lossy() {
		self.json = ctx.Dom.CopyJsonString()
		ctx.Json = self.json
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

	dec, err := findOrCompile(etp)
	if err != nil {
		return err
	}

	err = dec.FromDom(vp, ctx.Dom.Root(), &ctx)
	if err != nil {
		return error_parse_internal(err, self.json)
	}
	return err
}
