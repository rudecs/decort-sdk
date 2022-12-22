package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list of all computes with their relationships
type AffinityGroupComputesRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Affinity group label
	// Required: true
	AffinityGroup string `url:"affinityGroup"`
}

func (rgrq AffinityGroupComputesRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}
	if rgrq.AffinityGroup == "" {
		return errors.New("field AffinityGroup cat not be empty")
	}

	return nil
}

// AffinityGroupComputes gets list of all computes with their relationships to another computes
func (r RG) AffinityGroupComputes(ctx context.Context, req AffinityGroupComputesRequest) (ListAffinityGroups, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/affinityGroupComputes"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAffinityGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
