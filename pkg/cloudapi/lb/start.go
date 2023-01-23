package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start load balancer
type StartRequest struct {
	// ID of the load balancer instance to start
	// Required: true
	LBID uint64 `url:"lbId"`
}

func (lbrq StartRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

// Start starts specified load balancer instance
func (l LB) Start(ctx context.Context, req StartRequest) (bool, error) {
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
