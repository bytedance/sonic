package decoder

import (
	"github.com/bytedance/sonic/dev/internal"
)

type context struct {
	internal.Context
	options Options
}

func newCtx(s string, opt Options) (context, error) {
	ctx, err := internal.NewContext(s, uint64(opt))
	if err != nil {
		return context{}, err
	}

	return context{
		Context: ctx,
		options: opt,
	}, nil
}
