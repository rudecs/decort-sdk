package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
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

func (r RG) AffinityGroupsGet(ctx context.Context, req AffinityGroupsGetRequest, options ...opts.DecortOpts) ([]uint64, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/affinityGroupsGet"
	agListRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	agList := []uint64{}
	if err := json.Unmarshal(agListRaw, &agList); err != nil {
		return nil, err
	}

	return agList, nil
}
