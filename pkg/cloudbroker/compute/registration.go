package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set compute registered in RT
type RegistrationRequest struct {
	// ID of the Compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Unique compute registration key
	// Required: true
	RegistrationKey string `url:"registrationKey" json:"registrationKey"`
}

func (crq RegistrationRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.RegistrationKey == "" {
		return errors.New("validation-error: field RegistrationKey must be set")
	}

	return nil
}

// Registration sets compute registered in RT
func (c Compute) Registration(ctx context.Context, req RegistrationRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/registration"

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
