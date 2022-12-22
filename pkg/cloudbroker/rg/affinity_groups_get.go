package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list computes from affinity group
type AffinityGroupsGetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Label affinity group
	// Required: true
	AffinityGroup string `url:"affinityGroup"`
}

func (rgrq AffinityGroupsGetRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if rgrq.AffinityGroup == "" {
		return errors.New("validation-error: field AffinityGroup must be set")
	}

	return nil
}

// AffinityGroupsGet gets list computes in the specified affinity group
func (r RG) AffinityGroupsGet(ctx context.Context, req AffinityGroupsGetRequest) ([]uint64, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/affinityGroupsGet"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]uint64, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
