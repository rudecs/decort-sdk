package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AffinityGroupComputesRequest struct {
	RGID          uint64 `url:"rgId"`
	AffinityGroup string `url:"affinityGroup"`
}

func (rgrq AffinityGroupComputesRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	if rgrq.AffinityGroup == "" {
		return errors.New("field AffinityGroup cat not be empty")
	}

	return nil
}

func (r RG) AffinityGroupComputes(ctx context.Context, req AffinityGroupComputesRequest, options ...opts.DecortOpts) (AffinityGroupComputeList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/affinityGroupComputes"
	agcListRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	agcList := AffinityGroupComputeList{}
	if err := json.Unmarshal(agcListRaw, &agcList); err != nil {
		return nil, err
	}

	return agcList, nil
}
