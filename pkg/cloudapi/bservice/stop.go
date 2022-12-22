package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for stop service
type StopRequest struct {
	// ID of the service to stop
	// Required: true
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq StopRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Stop stops service.
// Stopping a service technically means stopping computes from
// all service groups
func (b BService) Stop(ctx context.Context, req StopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/stop"

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
