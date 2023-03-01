package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for rename disk
type RenameRequest struct {
	// ID of the disk to rename
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId"`

	// New name of disk
	// Required: true
	Name string `url:"name" json:"name"`
}

func (drq RenameRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// Rename rename disk
func (d Disks) Rename(ctx context.Context, req RenameRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/rename"

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
