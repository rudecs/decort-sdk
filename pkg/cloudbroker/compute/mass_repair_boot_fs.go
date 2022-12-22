package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for repair boot disk filesystem on several computes
type MassRepairBootFSRequest struct {
	// IDs of compute instances which boot file systems will be repaired
	// Required: true
	ComputeIDs []uint64 `url:"computeIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq MassRepairBootFSRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// MassRepairBootFS repair boot disk filesystem on several computes
func (c Compute) MassRepairBootFS(ctx context.Context, req MassRepairBootFSRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/massRepairBootFs"

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
