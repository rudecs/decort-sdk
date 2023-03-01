package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable service
type DisableRequest struct {
	// ID of the service to disable
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`
}

func (bsrq DisableRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Disable disables service.
// Disabling a service technically means setting model status
// of all computes and service itself to DISABLED and stopping all computes.
func (b BService) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/delete"

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
