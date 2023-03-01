package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for reboot compute
type RebootRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq RebootRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field computeID must be set")
	}

	return nil
}

// Reboot reboot compute
func (c Compute) Reboot(ctx context.Context, req RebootRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/reboot"

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
