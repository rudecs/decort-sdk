package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AffinityGroupsGetRequest struct {
	RGID          uint64 `url:"rgId"`
	AffinityGroup string `url:"affinityGroup"`
}

func (rgrq AffinityGroupsGetRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	if rgrq.AffinityGroup == "" {
		return errors.New("validation-error: field AffinityGroup must be set")
	}

	return nil
}

func (r RG) AffinityGroupsGet(ctx context.Context, req AffinityGroupsGetRequest) ([]uint64, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/affinityGroupsGet"

	agListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	agList := make([]uint64, 0)

	err = json.Unmarshal(agListRaw, &agList)
	if err != nil {
		return nil, err
	}

	return agList, nil
}
