package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for detach and delete disk from compute
type DiskDelRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of disk instance
	// Required: true
	DiskID uint64 `url:"diskId"`

	// False if disk is to be deleted to recycle bin
	// Required: true
	Permanently bool `url:"permanently"`
}

func (crq DiskDelRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

// DiskDel delete disk and detach from compute
func (c Compute) DiskDel(ctx context.Context, req DiskDelRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/diskDel"

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
