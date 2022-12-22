package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for evict specified disk
type SnapshotEvictDiskRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of the disk instance
	// Required: true
	DiskID uint64 `url:"diskId"`
}

func (crq SnapshotEvictDiskRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

// SnapshotEvictDisk evict specified disk from all snapshots of a compute instance
func (c Compute) SnapshotEvictDisk(ctx context.Context, req SnapshotEvictDiskRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/snapshotEvictDisk"

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
