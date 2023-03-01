package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start service
type StartRequest struct {
	// ID of the service to start
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`
}

func (bsrq StartRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Start starts service.
// Starting a service technically means starting computes from all
// service groups according to group relations
func (b BService) Start(ctx context.Context, req StartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/start"

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
