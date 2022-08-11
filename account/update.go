package account

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type UpdateRequest struct {
	AccountId              uint64 `url:"accountId"`
	Name                   string `url:"name,omitempty"`
	MaxMemoryCapacity      uint   `url:"maxMemoryCapacity,omitempty"`
	MaxVDiskCapacity       uint   `url:"maxVDiskCapacity,omitempty"`
	MaxCPUCapacity         uint   `url:"maxCPUCapacity,omitempty"`
	MaxNetworkPeerTransfer uint   `url:"maxNetworkPeerTransfer,omitempty"`
	MaxNumPublicIP         uint   `url:"maxNumPublicIP,omitempty"`
	SendAccessEmails       bool   `url:"sendAccessEmails,omitempty"`
	GpuUnits               uint   `url:"gpu_units,omitempty"`
}

func (arq UpdateRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) Update(ctx context.Context, req UpdateRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/account/update"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := a.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
