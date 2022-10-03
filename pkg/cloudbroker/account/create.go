package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	Name                   string `url:"name"`
	Username               string `url:"username"`
	EmailAddress           string `url:"emailaddress,omitempty"`
	MaxMemoryCapacity      uint   `url:"maxMemoryCapacity,omitempty"`
	MaxVDiskCapacity       uint   `url:"maxVDiskCapacity,omitempty"`
	MaxCPUCapacity         uint   `url:"maxCPUCapacity,omitempty"`
	MaxNetworkPeerTransfer uint   `url:"maxNetworkPeerTransfer,omitempty"`
	MaxNumPublicIP         uint   `url:"maxNumPublicIP,omitempty"`
	SendAccessEmails       bool   `url:"sendAccessEmails,omitempty"`
	GPUUnits               uint   `url:"gpu_units,omitempty"`
}

func (arq CreateRequest) Validate() error {
	if arq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if arq.Username == "" {
		return errors.New("validation-error: field Username can not be empty")
	}
	return nil
}

func (a Account) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/account/create"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
