package extnet

import (
	"context"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

func (e Extnet) GetDefault(ctx context.Context, options ...opts.DecortOpts) (uint64, error) {
	url := "/extnet/getDefault"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := e.client.DecortApiCall(ctx, typed.POST, url, nil)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}
