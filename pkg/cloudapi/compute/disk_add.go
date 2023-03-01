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
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Name for disk
	// Required: true
	DiskName string `url:"diskName" json:"diskName"`

	// Disk size in GB
	// Required: true
	Size uint64 `url:"size" json:"size"`

	// Type of the disk
	// Should be one of:
	//	- D
	//	- B
	// Required: false
	DiskType string `url:"diskType,omitempty" json:"diskType,omitempty"`

	// Storage endpoint provider ID
	// By default the same with boot disk
	// Required: false
	SepID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool name
	// By default will be chosen automatically
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Optional description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Specify image id for create disk from template
	// Required: false
	ImageID uint64 `url:"imageId,omitempty" json:"imageId,omitempty"`
}

func (crq DiskAddRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Size == 0 {
		return errors.New("validation-error: field Size can not be empty or equal to 0")
	}
	if crq.DiskName == "" {
		return errors.New("validation-error: field DiskName can not be empty or equal to 0")
	}

	return nil
}

// DiskAdd creates new disk and attach to compute
func (c Compute) DiskAdd(ctx context.Context, req DiskAddRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/compute/diskAdd"

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
