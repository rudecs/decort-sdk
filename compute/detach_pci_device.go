package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DetachPciDeviceRequest struct {
	ComputeId uint64 `url:"computeId"`
	DeviceID  uint64 `url:"deviceId"`
}

func (crq DetachPciDeviceRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	if crq.DeviceID == 0 {
		return errors.New("validation-error: field DeviceID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) DetachPciDevice(ctx context.Context, req DetachPciDeviceRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/detachPciDevice"
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
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, nil
	}

	return result, nil
}
