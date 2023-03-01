package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restart load balancer
type RestartRequest struct {
	// ID of the load balancer instance to restart
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq RestartRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Restart restarts specified load balancer instance
func (lb LB) Restart(ctx context.Context, req RestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/restart"

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
