package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	AccountID              uint64 `url:"accountId"`
	Name                   string `url:"name,omitempty"`
	MaxMemoryCapacity      uint64 `url:"maxMemoryCapacity,omitempty"`
	MaxVDiskCapacity       uint64 `url:"maxVDiskCapacity,omitempty"`
	MaxCPUCapacity         uint64 `url:"maxCPUCapacity,omitempty"`
	MaxNetworkPeerTransfer uint64 `url:"maxNetworkPeerTransfer,omitempty"`
	MaxNumPublicIP         uint64 `url:"maxNumPublicIP,omitempty"`
	SendAccessEmails       bool   `url:"sendAccessEmails,omitempty"`
	GPUUnits               uint64 `url:"gpu_units,omitempty"`
}

func (arq UpdateRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/update"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
