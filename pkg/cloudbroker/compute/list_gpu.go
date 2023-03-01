package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list GPU for compute
type ListGPURequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Also list destroyed
	// Required: false
	ListDestroyed bool `url:"list_destroyed,omitempty" json:"list_destroyed,omitempty"`
}

func (crq ListGPURequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// ListVGPU gets list GPU for compute
func (c Compute) ListGPU(ctx context.Context, req ListGPURequest) ([]interface{}, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/listGpu"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]interface{}, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
