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
		return errors.New("field RGID can not be empty or equal to 0")
	}

	if rgrq.AffinityGroup == "" {
		return errors.New("field AffinityGroup cat not be empty")
	}

	return nil
}

func (r RG) AffinityGroupsGet(ctx context.Context, req AffinityGroupsGetRequest) ([]uint64, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/affinityGroupsGet"
	agListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	agList := []uint64{}
	if err := json.Unmarshal(agListRaw, &agList); err != nil {
		return nil, err
	}

	return agList, nil
}
