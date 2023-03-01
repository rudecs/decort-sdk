package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for detach vgpu for compute
type DetachGPURequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Identifier virtual GPU
	// Required: false
	VGPUID int64 `url:"vgpuId,omitempty" json:"vgpuId,omitempty"`
}

func (crq DetachGPURequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// DetachGPU detach vgpu for compute.
// If param vgpuid is equivalent -1, then detach all vgpu for compute
func (c Compute) DetachGPU(ctx context.Context, req DetachGPURequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/detachGpu"

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
