package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable load balancer
type DisableRequest struct {
	// ID of the load balancer instance to disable
	// Required: true
	LBID uint64 `url:"lbId"`
}

func (lbrq DisableRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Disable disables specified load balancer instance
func (lb LB) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/disable"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
