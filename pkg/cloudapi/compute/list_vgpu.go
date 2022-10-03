package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListVGPURequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq ListVGPURequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) ListVGPU(ctx context.Context, req ListVGPURequest) ([]interface{}, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/listVgpu"

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
