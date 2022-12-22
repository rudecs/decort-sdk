package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clone compute instance
type CloneRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Name of the clone
	// Required: true
	Name string `url:"name"`

	// Timestamp of the parent's snapshot to create clone from
	// Required: false
	SnapshotTimestamp uint64 `url:"snapshotTimestamp"`

	// Name of the parent's snapshot to create clone from
	// Required: false
	SnapshotName string `url:"snapshotName"`
}

func (crq CloneRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Name == "" {
		return errors.New("validation-error: field Name can not be empty or equal to 0")
	}

	return nil
}

// Clone clones compute instance
func (c Compute) Clone(ctx context.Context, req CloneRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/compute/clone"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
