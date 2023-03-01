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
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// ID of disk instance
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId"`

	// False if disk is to be deleted to recycle bin
	// Required: true
	Permanently bool `url:"permanently" json:"permanently"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq DiskDelRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

// DiskDel delete disk and detach from compute
func (c Compute) DiskDel(ctx context.Context, req DiskDelRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/diskDel"

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
