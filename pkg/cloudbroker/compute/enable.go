package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for enable compute
type EnableRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq EnableRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field computeID must be set")
	}

	return nil
}

// Enable enables compute
func (c Compute) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/enable"

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
