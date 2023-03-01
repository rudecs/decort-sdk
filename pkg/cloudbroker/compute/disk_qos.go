package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for change QOS of the disk
type DiskQOSRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// ID of the disk to apply limits
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId"`

	// Limit IO for a certain disk total and read/write options are not allowed to be combined
	// Required: true
	Limits string `url:"limits" json:"limits"`
}

func (crq DiskQOSRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if crq.Limits == "" {
		return errors.New("validation-error: field Limits must be set")
	}

	return nil
}

// DiskQOS change QOS of the disk
func (c Compute) DiskQOS(ctx context.Context, req DiskQOSRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/diskQos"

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
