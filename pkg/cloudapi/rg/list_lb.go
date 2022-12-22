package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list load balancers
type ListLBRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`
}

func (rgrq ListLBRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

// ListLB gets list all load balancers in the specified resource group, accessible by the user
func (r RG) ListLB(ctx context.Context, req ListLBRequest) (ListLB, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listLb"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListLB{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
