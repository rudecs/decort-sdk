package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete snapshot
type SnapshotDeleteRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Text label of snapshot to delete
	// Required: true
	Label string `url:"label"`
}

func (crq SnapshotDeleteRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Label == "" {
		return errors.New("validation-error: field Label can not be empty")
	}

	return nil
}

// SnapshotDelete delete specified compute snapshot
func (c Compute) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/snapshotDelete"

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
