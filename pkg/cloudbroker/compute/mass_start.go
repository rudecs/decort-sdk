package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start several computes
type MassStartRequest struct {
	// IDs of compute instances to start
	// Required: true
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq MassStartRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// MassStart starts jobs to start several computes
func (c Compute) MassStart(ctx context.Context, req MassStartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/massStart"

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
