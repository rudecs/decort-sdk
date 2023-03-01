package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for resize disk
type ResizeRequest struct {
	// ID of the disk to resize
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId"`

	// New size of the disk in GB
	// Required: true
	Size uint64 `url:"size" json:"size"`
}

func (drq ResizeRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Size == 0 {
		return errors.New("validation-error: field Size must be set")
	}
	return nil
}

// Resize resize disk
// Returns 200 if disk is resized online, else will return 202,
// in that case please stop and start your machine after changing the disk size, for your changes to be reflected.
// This method will not be used for disks, assigned to computes. Only unassigned disks and disks, assigned with "old" virtual machines.
func (d Disks) Resize(ctx context.Context, req ResizeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/resize"

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

// Resize2 resize disk
// Returns 200 if disk is resized online, else will return 202,
// in that case please stop and start your machine after changing the disk size, for your changes to be reflected.
// This method will not be used for disks, assigned to "old" virtual machines. Only unassigned disks and disks, assigned with computes.
func (d Disks) Resize2(ctx context.Context, req ResizeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/resize2"

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
