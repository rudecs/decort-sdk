package vins

import (
	"context"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type IPReserveRequest struct {
	VinsId    uint64 `url:"vinsId"`
	Type      string `url:"type"`
	IPAddr    string `url:"ipAddr,omitempty"`
	MAC       string `url:"mac,omitempty"`
	ComputeId uint64 `url:"computeId,omitempty"`
}

func (vrq IPReserveRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	if vrq.Type == "" {
		return errors.New("validation-error: field Type can not be empty")
	}

	return nil
}

func (v Vins) IPReserve(ctx context.Context, req IPReserveRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/vins/ipReserve"
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
		return "", err
	}

	return string(res), nil

}
