package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for reboot several computes
type MassRebootRequest struct {
	// IDs of compute instances to reboot
	// Required: true
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq MassRebootRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// MassReboot starts jobs to reboot several computes
func (c Compute) MassReboot(ctx context.Context, req MassRebootRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/massReboot"

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
