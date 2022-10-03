package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

func (r RG) ListLB(ctx context.Context, req ListLBRequest) (LBList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listLb"
	lbListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	lbList := LBList{}
	if err := json.Unmarshal(lbListRaw, &lbList); err != nil {
		return nil, err
	}

	return lbList, nil
}
