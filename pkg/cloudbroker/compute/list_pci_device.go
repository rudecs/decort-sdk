package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list PCI devices
type ListPCIDeviceRequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq ListPCIDeviceRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// ListPCIDevice gets list PCI device
func (c Compute) ListPCIDevice(ctx context.Context, req ListPCIDeviceRequest) (ListPCIDevices, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/listPciDevice"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListPCIDevices{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
