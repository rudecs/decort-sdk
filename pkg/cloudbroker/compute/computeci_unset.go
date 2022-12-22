package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for unset compite CI
type ComputeCIUnsetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq ComputeCIUnsetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// ComputeCIUnset unsets compute CI ID from compute
func (c Compute) ComputeCIUnset(ctx context.Context, req ComputeCIUnsetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/computeciUnset"

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
