package vins

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateInAccountRequest struct {
	Name               string `url:"name"`
	AccountId          uint64 `url:"accountId"`
	GID                uint64 `url:"gid,omitempty"`
	IPCidr             string `url:"ipcidr,omitempty"`
	Description        string `url:"desc,omitempty"`
	PreReservationsNum uint   `url:"preReservationsNum,omitempty"`
}

func (vrq CreateInAccountRequest) Validate() error {
	if vrq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if vrq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) CreateInAccount(ctx context.Context, req CreateInAccountRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/vins/createInAccount"
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