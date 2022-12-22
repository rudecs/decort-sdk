package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete bind
type FrontendBindDeleteRequest struct {
	// ID of the load balancer instance to FrontendBindDelete
	// Required: true
	LBID uint64 `url:"lbId"`

	// Name of the frontend to delete
	// Required: true
	FrontendName string `url:"frontendName"`

	// Name of the binding to delete
	// Required: true
	BindingName string `url:"bindingName"`
}

func (lbrq FrontendBindDeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}
	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName must be set")
	}
	if lbrq.BindingName == "" {
		return errors.New("validation-error: field BindingName must be set")
	}

	return nil
}

// FrontendBindDelete deletes binding from the specified load balancer frontend
func (lb LB) FrontendBindDelete(ctx context.Context, req FrontendBindDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/frontendBindDelete"

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
