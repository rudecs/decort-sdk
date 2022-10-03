package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type AttachPciDeviceRequest struct {
	ComputeID uint64 `url:"computeId"`
	DeviceID  uint64 `url:"deviceId"`
}

func (crq AttachPciDeviceRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.DeviceID == 0 {
		return errors.New("validation-error: field DeviceID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) AttachPciDevice(ctx context.Context, req AttachPciDeviceRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/attachPciDevice"

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
