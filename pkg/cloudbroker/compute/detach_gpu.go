package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for detach VGPU for compute
type DetachGPURequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Identifier virtual GPU
	// Required: false
	VGPUID int64 `url:"vgpuId,omitempty"`
}

func (crq DetachGPURequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// DetachGPU detach VGPU for compute.
// If param VGPU ID is equivalent -1, then detach all VGPU for compute
func (c Compute) DetachGPU(ctx context.Context, req DetachGPURequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/detachGpu"

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
