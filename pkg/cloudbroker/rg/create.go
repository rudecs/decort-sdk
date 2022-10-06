package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	AccountID              uint64 `url:"accountId"`
	GID                    uint64 `url:"gid"`
	Name                   string `url:"name"`
	MaxMemoryCapacity      uint64 `url:"maxMemoryCapacity,omitempty"`
	MaxVDiskCapacity       uint64 `url:"maxVDiskCapacity,omitempty"`
	MaxCPUCapacity         uint64 `url:"maxCPUCapacity,omitempty"`
	MaxNetworkPeerTransfer uint64 `url:"maxNetworkPeerTransfer,omitempty"`
	MaxNumPublicIP         uint64 `url:"maxNetworkPeerTransfer,omitempty"`
	Owner                  string `url:"owner,omitempty"`
	DefNet                 string `url:"def_net,omitempty"`
	IPCIDR                 string `url:"ipcidr,omitempty"`
	Desc                   string `url:"desc,omitempty"`
	Reason                 string `url:"reason,omitempty"`
	ExtNetID               uint64 `url:"extNetId,omitempty"`
	ExtIP                  string `url:"extIp,omitempty"`
	RegisterComputes       bool   `url:"registerComputes,omitempty"`
}

func (rgrq CreateRequest) Validate() error {
	if rgrq.AccountID == 0 {
		return errors.New("field AccountID can not be empty or equal to 0")
	}

	if rgrq.GID == 0 {
		return errors.New("field GID can not be empty or equal to 0")
	}

	if len(rgrq.Name) < 2 {
		return errors.New("field Name can not be shorter than two bytes")
	}

	return nil
}

func (r RG) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/rg/create"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
