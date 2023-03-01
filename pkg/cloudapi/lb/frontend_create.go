package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create frontend
type FrontendCreateRequest struct {
	// ID of the load balancer instance to FrontendCreate
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`

	// Must be unique among all frontends of
	// this load balancer - name of the new frontend to create
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName"`

	// Should be one of the backends existing on
	// this load balancer - name of the backend to use
	// Required: true
	BackendName string `url:"backendName" json:"backendName"`
}

func (lbrq FrontendCreateRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName can not be empty")
	}
	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName can not be empty")
	}

	return nil
}

// FrontendCreate creates new frontend on the specified load balancer
func (l LB) FrontendCreate(ctx context.Context, req FrontendCreateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/frontendCreate"

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
