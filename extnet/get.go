package extnet

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	NetId uint64 `url:"net_id"`
}

func (erq GetRequest) Validate() error {
	if erq.NetId == 0 {
		return errors.New("validation-error: field NetId can not be empty or equal to 0")
	}
	return nil
}

func (e Extnet) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*ExtnetDetailed, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/extnet/get"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	extnetRaw, err := e.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	extnet := &ExtnetDetailed{}
	err = json.Unmarshal(extnetRaw, &extnet)
	if err != nil {
		return nil, err
	}

	return extnet, nil

}
