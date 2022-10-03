package computeci

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	ComputeCIID uint64 `url:"computeciId"`
}

func (krq GetRequest) Validate() error {
	if krq.ComputeCIID == 0 {
		return errors.New("field ComputeCIID can not be empty or equal to 0")
	}

	return nil
}

func (c ComputeCI) Get(ctx context.Context, req GetRequest) (*ComputeCIRecord, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/computeci/get"
	computeciRaw, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	computeci := &ComputeCIRecord{}
	if err := json.Unmarshal(computeciRaw, computeci); err != nil {
		return nil, err
	}

	return computeci, nil
}
