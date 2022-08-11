package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AttachGPURequest struct {
	ComputeId uint64 `url:"computeId"`
	VGPUID    uint64 `url:"vgpuId"`
}

func (crq AttachGPURequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	if crq.VGPUID == 0 {
		return errors.New("validation-error: field VGPUID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) AttachGPU(ctx context.Context, req AttachGPURequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/compute/attachGpu"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, nil
	}

	return result, nil
}
