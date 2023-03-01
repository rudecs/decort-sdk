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
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

// Restart restarts specified load balancer instance
func (l LB) Restart(ctx context.Context, req RestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/restart"

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
