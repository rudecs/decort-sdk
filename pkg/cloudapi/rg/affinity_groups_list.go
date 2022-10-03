package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AffinityGroupsListRequest struct {
	RGID uint64 `url:"rgId"`
}

func (rgrq AffinityGroupsListRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) AffinityGroupsList(ctx context.Context, req AffinityGroupsListRequest) (map[string][]uint64, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/affinityGroupsList"
	agListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	agList := map[string][]uint64{}
	if err := json.Unmarshal(agListRaw, &agList); err != nil {
		return nil, err
	}

	return agList, nil
}
