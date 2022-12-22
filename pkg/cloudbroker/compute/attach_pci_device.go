package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for atttach PCI device
type AttachPCIDeviceRequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// PCI device ID
	// Required: true
	DeviceID uint64 `url:"deviceId"`
}

func (crq AttachPCIDeviceRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.DeviceID == 0 {
		return errors.New("validation-error: field DeviceID must be set")
	}

	return nil
}

// AttachPCIDevice attach PCI device
func (c Compute) AttachPCIDevice(ctx context.Context, req AttachPCIDeviceRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/attachPciDevice"

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
