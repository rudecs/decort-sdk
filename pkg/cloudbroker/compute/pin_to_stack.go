package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for pin comptute to stack
type PinToStackRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Stack ID to pin to
	// Required: true
	TargetStackID uint64 `url:"targetStackId" json:"targetStackId"`

	// Try to migrate or not if compute in running states
	// Required: false
	Force bool `url:"force" json:"force"`
}

func (crq PinToStackRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.TargetStackID == 0 {
		return errors.New("validation-error: field TargetStackID must be set")
	}

	return nil
}

// PinToStack pin compute to current stack
func (c Compute) PinToStack(ctx context.Context, req PinToStackRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/compute/pinToStack"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
