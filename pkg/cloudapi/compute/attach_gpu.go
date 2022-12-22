package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for attach GPU for compute
type AttachGPURequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Identifier vGPU
	// Required: true
	VGPUID uint64 `url:"vgpuId"`
}

func (crq AttachGPURequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.VGPUID == 0 {
		return errors.New("validation-error: field VGPUID can not be empty or equal to 0")
	}

	return nil
}

// AttachGPU attach GPU for compute, returns vgpu id on success
func (c Compute) AttachGPU(ctx context.Context, req AttachGPURequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/compute/attachGpu"

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
