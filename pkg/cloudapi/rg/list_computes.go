package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

func (r RG) ListComputes(ctx context.Context, req ListComputesRequest) (ComputeList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listComputes"
	computeListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	computeList := ComputeList{}
	if err := json.Unmarshal(computeListRaw, &computeList); err != nil {
		return nil, err
	}

	return computeList, nil
}
