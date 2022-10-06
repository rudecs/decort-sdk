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
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) ListLB(ctx context.Context, req ListLBRequest) (ListLB, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/listLb"

	lbListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	lbList := ListLB{}

	err = json.Unmarshal(lbListRaw, &lbList)
	if err != nil {
		return nil, err
	}

	return lbList, nil
}
