package computeci

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
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

func (c ComputeCI) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*ComputeCIRecord, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/computeci/get"
	computeciRaw, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	computeci := &ComputeCIRecord{}
	if err := json.Unmarshal(computeciRaw, computeci); err != nil {
		return nil, err
	}

	return computeci, nil
}
