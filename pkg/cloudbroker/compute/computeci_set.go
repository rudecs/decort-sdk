package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set compute CI
type ComputeCISetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of the Compute CI
	// Required: true
	ComputeCIID uint64 `url:"computeciId"`
}

func (crq ComputeCISetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.ComputeCIID == 0 {
		return errors.New("validation-error: field ComputeCIID must be set")
	}

	return nil
}

// ComputeCISet sets compute CI ID for compute
func (c Compute) ComputeCISet(ctx context.Context, req ComputeCISetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/computeciSet"

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
