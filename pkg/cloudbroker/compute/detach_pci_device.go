package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for detach PCI device
type DetachPCIDeviceRequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// PCI device ID
	// Required: true
	DeviceID uint64 `url:"deviceId"`
}

func (crq DetachPCIDeviceRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DeviceID == 0 {
		return errors.New("validation-error: field DeviceID must be set")
	}

	return nil
}

// DetachPCIDevice detach PCI device
func (c Compute) DetachPciDevice(ctx context.Context, req DetachPCIDeviceRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/detachPciDevice"

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
