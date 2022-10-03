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
	MaxNumPublicIP         uint64 `url:"maxNumPublicIP,omitempty"`
	RegisterComputes       bool   `url:"registerComputes,omitempty"`
	Reason                 string `url:"reason,omitempty"`
}

func (rgrq UpdateRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/rg/update"
	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
