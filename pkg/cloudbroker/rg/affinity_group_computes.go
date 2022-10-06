package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AffinityGroupComputesRequest struct {
	RGID          uint64 `url:"rgId"`
	AffinityGroup string `url:"affinityGroup"`
}

func (rgrq AffinityGroupComputesRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	if rgrq.AffinityGroup == "" {
		return errors.New("validation-error: field AffinityGroup must be set")
	}

	return nil
}

func (r RG) AffinityGroupComputes(ctx context.Context, req AffinityGroupComputesRequest) (AffinityGroupComputeList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/affinityGroupComputes"

	agcListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	agcList := AffinityGroupComputeList{}

	if err := json.Unmarshal(agcListRaw, &agcList); err != nil {
		return nil, err
	}

	return agcList, nil
}
