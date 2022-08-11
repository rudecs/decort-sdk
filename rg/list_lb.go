package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListLBRequest struct {
	RGID uint64 `url:"rgId"`
}

func (rgrq ListLBRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) ListLB(ctx context.Context, req ListLBRequest, options ...opts.DecortOpts) (LBList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listLb"
	lbListRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	lbList := LBList{}
	if err := json.Unmarshal(lbListRaw, &lbList); err != nil {
		return nil, err
	}

	return lbList, nil
}
