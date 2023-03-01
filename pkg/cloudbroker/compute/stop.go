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

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq StopRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field computeID must be set")
	}

	return nil
}

// Stop stops compute
func (c Compute) Stop(ctx context.Context, req StopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/stop"

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
