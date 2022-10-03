package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListPCIDeviceRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq ListPCIDeviceRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) ListPCIDevice(ctx context.Context, req ListPCIDeviceRequest) ([]interface{}, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/listPciDevice"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
