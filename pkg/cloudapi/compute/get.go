package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq GetRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Get(ctx context.Context, req GetRequest) (*ComputeRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/get"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
