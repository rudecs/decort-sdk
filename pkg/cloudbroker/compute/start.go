package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start compute
type StartRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of CD-ROM live image to boot
	// Required: false
	AltBootID uint64 `url:"altBootId,omitempty"`

	// ID of stack to start compute
	// Required: false
	StackID uint64 `url:"stackId,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq StartRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field computeID must be set")
	}

	return nil
}

// Start starts compute
func (c Compute) Start(ctx context.Context, req StartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/start"

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
