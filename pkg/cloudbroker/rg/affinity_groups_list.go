package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list of affinity groups from resource group
type AffinityGroupsListRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`
}

func (rgrq AffinityGroupsListRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

// AffinityGroupsList gets all currently defined affinity groups in this resource group with compute IDs
func (r RG) AffinityGroupsList(ctx context.Context, req AffinityGroupsListRequest) (map[string][]uint64, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/affinityGroupsList"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make(map[string][]uint64)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
