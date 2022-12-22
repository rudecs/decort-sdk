package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for several stop computes
type MassStopRequest struct {
	// IDs of compute instances to stop
	// Required: true
	ComputeIDs []uint64 `url:"computeIds"`

	// Force stop compute
	// Required: false
	Force bool `url:"force,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq MassStopRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// MassStop starts jobs to stop several computes
func (c Compute) MassStop(ctx context.Context, req MassStopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/massStop"

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
