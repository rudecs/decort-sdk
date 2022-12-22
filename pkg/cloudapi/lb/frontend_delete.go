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
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName can not be empty")
	}

	return nil
}

// FrontendDelete deletes frontend from the specified load balancer.
// Warning: you cannot undo this action!
func (l LB) FrontendDelete(ctx context.Context, req FrontendDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/frontendDelete"

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
