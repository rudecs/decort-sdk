package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for stop load balancer
type StopRequest struct {
	// ID of the LB instance to stop
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq StopRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Stop stops specified load balancer instance
func (lb LB) Stop(ctx context.Context, req StopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/stop"

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
