package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restore a deleted unattached disk
type RestoreRequest struct {
	// ID of the disk to restore
	// Required: true
	DiskID uint64 `url:"diskId"`

	// Reason for restoring the disk
	// Required: true
	Reason string `url:"reason"`
}

func (drq RestoreRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// Restore restore a deleted unattached disk from recycle bin
func (d Disks) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/restore"

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
