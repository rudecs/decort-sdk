package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for change disk size
type DiskResizeRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of the disk to resize
	// Required: true
	DiskID uint64 `url:"diskId"`

	// New disk size
	// Required: true
	Size uint64 `url:"size"`
}

func (crq DiskResizeRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}
	if crq.Size == 0 {
		return errors.New("validation-error: field Size can not be empty or equal to 0")
	}

	return nil
}

// DiskResize change disk size
func (c Compute) DiskResize(ctx context.Context, req DiskResizeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/diskResize"

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
