package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListVGPURequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq ListVGPURequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) ListVGPU(ctx context.Context, req ListVGPURequest, options ...opts.DecortOpts) ([]interface{}, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/listVgpu"
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
		return nil, err
	}

	pciDeviceList := []interface{}{}

	err = json.Unmarshal(res, &pciDeviceList)
	if err != nil {
		return nil, err
	}

	return pciDeviceList, nil
}