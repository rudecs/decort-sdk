package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for rollback snapshot
type SnapshotRollbackRequest struct {
	// ID of the disk
	// Required: true
	DiskID uint64 `url:"diskId"`

	// Label of the snapshot to rollback
	// Required: true
	Label string `url:"label"`

	// Timestamp of the snapshot to rollback
	// Required: true
	TimeStamp uint64 `url:"timestamp"`
}

func (drq SnapshotRollbackRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Label == "" {
		return errors.New("validation-error: field Label must be set")
	}
	if drq.TimeStamp == 0 {
		return errors.New("validation-error: field TimeStamp must be set")
	}

	return nil
}

// SnapshotRollback rollback an individual disk snapshot
func (d Disks) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/snapshotRollback"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
