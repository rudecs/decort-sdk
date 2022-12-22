package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restore load balancer
type RestoreRequest struct {
	// ID of the load balancer instance to restore
	// Required: true
	LBID uint64 `url:"lbId"`
}

func (lbrq RestoreRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Restore restore load balancer from recycle bin
func (lb LB) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/restore"

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
