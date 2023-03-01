package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for enable load balancer
type EnableRequest struct {
	// ID of the load balancer instance to enable
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq EnableRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Enable enables specified load balancer instance
func (lb LB) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/enable"

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
