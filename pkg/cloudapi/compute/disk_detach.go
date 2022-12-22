package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for detach disk from compute
type DiskDetachRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of the disk to detach
	// Required: true
	DiskID uint64 `url:"diskId"`
}

func (crq DiskDetachRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

// DiskDetach detach disk from compute
func (c Compute) DiskDetach(ctx context.Context, req DiskDetachRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/diskDetach"

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
