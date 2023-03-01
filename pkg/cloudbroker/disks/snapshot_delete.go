package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete snapshot
type SnapshotDeleteRequest struct {
	// ID of disk to delete
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId"`

	// Label of the snapshot to delete
	// Required: false
	Label string `url:"label" json:"label"`
}

func (drq SnapshotDeleteRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Label == "" {
		return errors.New("validation-error: field Label must be set")
	}

	return nil
}

// SnapshotDelete deletes a snapshot
func (d Disks) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/snapshotDelete"

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
