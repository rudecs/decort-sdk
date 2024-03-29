package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for attach disk to compute
type DiskAttachRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// ID of the disk to attach
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq DiskAttachRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

// DiskAttach attach disk to compute
func (c Compute) DiskAttach(ctx context.Context, req DiskAttachRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/diskAttach"

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
