package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete backend
type BackendDeleteRequest struct {
	// ID of the load balancer instance to BackendDelete
	// Required: true
	LBID uint64 `url:"lbId"`

	// Cannot be emtpy string - name of the backend to delete
	// Required: true
	BackendName string `url:"backendName"`
}

func (lbrq BackendDeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}
	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName must be set")
	}

	return nil
}

// BackendDelete deletes backend from the specified load balancer.
// Warning: you cannot undo this action!
func (lb LB) BackendDelete(ctx context.Context, req BackendDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/backendDelete"

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
