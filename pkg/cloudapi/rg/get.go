package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq GetRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) Get(ctx context.Context, req GetRequest) (*ResourceGroup, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/get"
	rgRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	rg := &ResourceGroup{}
	if err := json.Unmarshal(rgRaw, rg); err != nil {
		return nil, err
	}

	return rg, nil
}
