package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create and attach disk to compute
type DiskAddRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Name for disk
	// Required: true
	DiskName string `url:"diskName"`

	// Disk size in GB
	// Required: true
	Size uint64 `url:"size"`

	// Type of the disk
	// Should be one of:
	//	- D
	//	- B
	// Required: false
	DiskType string `url:"diskType,omitempty"`

	// Storage endpoint provider ID
	// By default the same with boot disk
	// Required: false
	SepID uint64 `url:"sepId,omitempty"`

	// Pool name
	// By default will be chosen automatically
	// Required: false
	Pool string `url:"pool,omitempty"`

	// Optional description
	// Required: false
	Description string `url:"desc,omitempty"`

	// Specify image id for create disk from template
	// Required: false
	ImageID uint64 `url:"imageId,omitempty"`
}

func (crq DiskAddRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DiskName == "" {
		return errors.New("validation-error: field DiskName must be set")
	}
	if crq.Size == 0 {
		return errors.New("validation-error: field Size must be set")
	}

	return nil
}

// DiskAdd creates new disk and attach to compute
func (c Compute) DiskAdd(ctx context.Context, req DiskAddRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/compute/diskAdd"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
