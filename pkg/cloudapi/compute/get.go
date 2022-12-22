package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request for get information about compute
type GetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq GetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Get Gets information about compute
func (c Compute) Get(ctx context.Context, req GetRequest) (*RecordCompute, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/get"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordCompute{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
