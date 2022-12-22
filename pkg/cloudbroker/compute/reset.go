package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for reset compute
type ResetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq ResetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// Reset reset compute
func (c Compute) Reset(ctx context.Context, req ResetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/reset"

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
