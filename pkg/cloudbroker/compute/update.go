package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update compute
type UpdateRequest struct {
	// ID of the compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// New name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// New description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq UpdateRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// Update updates some properties of the compute
func (c Compute) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/update"

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
