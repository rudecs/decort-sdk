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
	ComputeID uint64 `url:"computeId"`
}

func (crq ListPCIDeviceRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// ListPCIDevice gets list PCI device
func (c Compute) ListPCIDevice(ctx context.Context, req ListPCIDeviceRequest) ([]interface{}, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/listPciDevice"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := []interface{}{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
