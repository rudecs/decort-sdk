package vins

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateInRGRequest struct {
	Name               string `url:"name"`
	RGID               uint64 `url:"rgId"`
	IPCidr             string `url:"ipcidr,omitempty"`
	ExtNetId           uint64 `url:"extNetId,omitempty"`
	ExtIP              string `url:"extIp,omitempty"`
	Description        string `url:"desc,omitempty"`
	PreReservationsNum uint   `url:"preReservationsNum,omitempty"`
}

func (vrq CreateInRGRequest) Validate() error {
	if vrq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if vrq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) CreateInRG(ctx context.Context, req CreateInRGRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/vins/createInRG"
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
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
