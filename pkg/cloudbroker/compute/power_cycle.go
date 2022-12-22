package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for force stop and start compute
type PowerCycleRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq PowerCycleRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// PowerCycle makes force stop and start compute
func (c Compute) PowerCycle(ctx context.Context, req PowerCycleRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/powerCycle"

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
