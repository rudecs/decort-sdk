package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for multiple disks
type DeleteDisksRequest struct {
	// List of disk ids to delete
	// Required: true
	DisksIDs []uint64 `url:"diskIds"`

	// Reason for deleting the disks
	// Required: true
	Reason string `url:"reason"`

	// Whether to completely delete the disks, works only with non attached disks
	// Required: false
	Permanently bool `url:"permanently,omitempty"`
}

func (drq DeleteDisksRequest) validate() error {
	if len(drq.DisksIDs) == 0 {
		return errors.New("validation-error: field DiskIDs must be set")
	}
	if drq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// DeleteDisks deletes multiple disks permanently
func (d Disks) DeleteDisks(ctx context.Context, req DeleteDisksRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/deleteDisks"

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
