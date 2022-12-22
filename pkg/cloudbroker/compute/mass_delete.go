package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete several computes
type MassDeleteRequest struct {
	// IDs of compute instances to delete
	// Required: true
	ComputeIDs []uint64 `url:"computeIds"`

	// Delete computes permanently
	// Required: false
	Permanently bool `url:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq MassDeleteRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// MassDelete starts jobs to delete several computes
func (c Compute) MassDelete(ctx context.Context, req MassDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/massDelete"

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
