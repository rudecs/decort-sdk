package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request for delete disk
type DeleteRequest struct {
	// ID of disk to delete
	// Required: true
	DiskID uint64 `url:"diskId"`

	// Detach disk from machine first
	// Required: false
	Detach bool `url:"detach,omitempty"`

	// Whether to completely delete the disk, works only with non attached disks
	// Required: false
	Permanently bool `url:"permanently,omitempty"`

	// Reason to delete
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (d DeleteRequest) validate() error {
	if d.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

// Delete deletes disk by ID
func (d Disks) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/delete"

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
