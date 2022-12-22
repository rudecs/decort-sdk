package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable service
type EnableRequest struct {
	// ID of the service to enable
	// Required: true
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq EnableRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Enable enables service.
// Enabling a service technically means setting model status of
// all computes and service itself to ENABLED.
// It does not start computes.
func (b BService) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/enable"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
