package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	RGID                   uint64 `url:"rgId"`
	Name                   string `url:"name,omitempty"`
	Desc                   string `url:"desc,omitempty"`
	MaxMemoryCapacity      uint64 `url:"maxMemoryCapacity,omitempty"`
	MaxVDiskCapacity       uint64 `url:"maxVDiskCapacity,omitempty"`
	MaxCPUCapacity         uint64 `url:"maxCPUCapacity,omitempty"`
	MaxNetworkPeerTransfer uint64 `url:"maxNetworkPeerTransfer,omitempty"`
	MaxNumPublicIP         uint64 `url:"maxNetworkPeerTransfer,omitempty"`
	RegisterComputes       bool   `url:"registerComputes,omitempty"`
	Reason                 string `url:"reason"`
}

func (rgrq UpdateRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/update"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
