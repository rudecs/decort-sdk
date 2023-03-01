package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for stop load balancer
type StopRequest struct {
	// ID of the load balancer instance to stop
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq StopRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

// Stop stops specified load balancer instance
func (l LB) Stop(ctx context.Context, req StopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/start"

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
