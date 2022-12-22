package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete frontend
type FrontendDeleteRequest struct {
	// ID of the load balancer instance to FrontendDelete
	// Required: true
	LBID uint64 `url:"lbId"`

	// Name of the frontend to delete
	// Required: true
	FrontendName string `url:"frontendName"`
}

func (lbrq FrontendDeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}
	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName must be set")
	}

	return nil
}

// FrontendDelete deletes frontend from the specified load balancer.
// Warning: you cannot undo this action!
func (lb LB) FrontendDelete(ctx context.Context, req FrontendDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/frontendDelete"

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
