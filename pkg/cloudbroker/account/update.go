package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	AccountID              uint64 `url:"accountId"`
	Name                   string `url:"name"`
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
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

func (a Account) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/update"

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
