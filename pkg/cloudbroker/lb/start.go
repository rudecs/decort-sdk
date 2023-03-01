package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start load balancer
type StartRequest struct {
	// ID of the LB instance to start
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq StartRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Start starts specified load balancer instance
func (lb LB) Start(ctx context.Context, req StartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/start"

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
