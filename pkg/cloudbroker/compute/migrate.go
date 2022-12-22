package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for migrate compute
type MigrateRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Particular Stack ID to migrate this compute to
	// Required: false
	TargetStackID uint64 `url:"targetStackId,omitempty"`

	// If live migration fails, destroy compute
	// on source node and recreate on the target
	// Required: false
	Force bool `url:"force,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq MigrateRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// Migrate migrates compute to another stack
func (c Compute) Migrate(ctx context.Context, req MigrateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/migrate"

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
