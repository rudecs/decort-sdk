package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for stop compute
type StopRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Force stop compute
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

func (crq StopRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Stop stops compute
func (c Compute) Stop(ctx context.Context, req StopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/stop"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
