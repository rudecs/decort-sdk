package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type AttachGPURequest struct {
	ComputeID uint64 `url:"computeId"`
	VGPUID    uint64 `url:"vgpuId"`
}

func (crq AttachGPURequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.VGPUID == 0 {
		return errors.New("validation-error: field VGPUID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) AttachGPU(ctx context.Context, req AttachGPURequest) (uint64, error) {
	err := req.Validate()
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
