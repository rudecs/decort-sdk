package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for rollback
type SnapshotRollbackRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Text label of snapshot to rollback
	// Required: true
	Label string `url:"label"`
}

func (crq SnapshotRollbackRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Label == "" {
		return errors.New("validation-error: field Label can not be empty")
	}

	return nil
}

// SnapshotRollback rollback specified compute snapshot
func (c Compute) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/snapshotRollback"

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
