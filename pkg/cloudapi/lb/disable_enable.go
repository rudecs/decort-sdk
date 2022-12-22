package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable/enable load balancer
type DisableEnableRequest struct {
	// ID of the load balancer instance to disable/enable
	// Required: true
	LBID uint64 `url:"lbId"`
}

func (lbrq DisableEnableRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

// Disable disables specified load balancer instance
func (l LB) Disable(ctx context.Context, req DisableEnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/disable"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}

// Enable enables specified load balancer instance
func (l LB) Enable(ctx context.Context, req DisableEnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/enable"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
