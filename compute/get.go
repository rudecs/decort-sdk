package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq GetRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*ComputeRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/get"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	compute := &ComputeRecord{}
	err = json.Unmarshal(res, compute)
	if err != nil {
		return nil, err
	}

	return compute, nil
}
