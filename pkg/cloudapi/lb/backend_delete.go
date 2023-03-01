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
	LBID uint64 `url:"lbId" json:"lbId"`

	// Cannot be emtpy string - name of the backend to delete
	// Required: true
	BackendName string `url:"backendName" json:"backendName"`
}

func (lbrq BackendDeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName can not be empty")
	}

	return nil
}

// BackendDelete deletes backend from the specified load balancer.
// Warning: you cannot undo this action!
func (l LB) BackendDelete(ctx context.Context, req BackendDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/backendDelete"

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
