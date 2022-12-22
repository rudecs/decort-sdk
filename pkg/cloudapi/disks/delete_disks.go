package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for multiple disks
type DisksDeleteRequest struct {
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

func (drq DisksDeleteRequest) validate() error {
	if len(drq.DisksIDs) == 0 {
		return errors.New("validation-error: field DisksIDs must include one or more disks ids")
	}
	if drq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// DeleteDisks deletes multiple disks permanently
func (d Disks) DeleteDisks(ctx context.Context, req DisksDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/deleteDisks"

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
