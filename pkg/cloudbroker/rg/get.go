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
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) Get(ctx context.Context, req GetRequest) (*ResourceGroup, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/get"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	getResult := ResourceGroup{}

	err = json.Unmarshal(res, &getResult)
	if err != nil {
		return nil, err
	}

	return &getResult, nil
}
