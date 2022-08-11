package vins

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type IPReleaseRequest struct {
	VinsId uint64 `url:"vinsId"`
	IPAddr string `url:"ipAddr,omitempty"`
	MAC    string `url:"mac,omitempty"`
}

func (vrq IPReleaseRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) IPRelese(ctx context.Context, req IPReleaseRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/vins/ipRelease"
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
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, nil
	}

	return result, nil

}
