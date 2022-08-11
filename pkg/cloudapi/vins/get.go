package vins

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	VinsId uint64 `url:"vinsId"`
}

func (vrq GetRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*VinsDetailed, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/vins/get"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := v.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	vins := &VinsDetailed{}

	err = json.Unmarshal(res, vins)
	if err != nil {
		return nil, err
	}

	return vins, nil

}
