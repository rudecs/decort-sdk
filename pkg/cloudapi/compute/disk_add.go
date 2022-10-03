package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DiskAddRequest struct {
	ComputeID   uint64 `url:"computeId"`
	DiskName    string `url:"diskName"`
	Size        uint64 `url:"size"`
	DiskType    string `url:"diskType,omitempty"`
	SepID       uint64 `url:"sepId,omitempty"`
	Pool        string `url:"pool,omitempty"`
	Description string `url:"desc,omitempty"`
	ImageID     uint64 `url:"imageId,omitempty"`
}

func (crq DiskAddRequest) Validate() error {
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

func (c Compute) DiskAdd(ctx context.Context, req DiskAddRequest) (uint64, error) {
	err := req.Validate()
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
