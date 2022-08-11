package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListComputesRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq ListComputesRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) ListComputes(ctx context.Context, req ListComputesRequest, options ...opts.DecortOpts) (ComputeList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listComputes"
	computeListRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	computeList := ComputeList{}
	if err := json.Unmarshal(computeListRaw, &computeList); err != nil {
		return nil, err
	}

	return computeList, nil
}
