package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for move compute new resource group
type MoveToRGRequest struct {
	// ID of the compute instance to move
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// ID of the target resource group
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// New name for the compute upon successful move,
	// if name change required.
	// Pass empty string if no name change necessary
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Should the compute be restarted upon successful move
	// Required: false
	Autostart bool `url:"autostart,omitempty" json:"autostart,omitempty"`

	// By default moving compute in a running state is not allowed.
	// Set this flag to True to force stop running compute instance prior to move.
	// Required: false
	ForceStop bool `url:"forceStop,omitempty" json:"forceStop,omitempty"`
}

func (crq MoveToRGRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

// MoveToRG moves compute instance to new resource group
func (c Compute) Validate(ctx context.Context, req MoveToRGRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/moveToRg"

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
