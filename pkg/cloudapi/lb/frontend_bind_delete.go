package lb

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for delete bind
type FrontendBindDeleteRequest struct {
	// ID of the load balancer instance to FrontendBindDelete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`

	// Name of the frontend to delete
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName"`

	// Name of the binding to delete
	// Required: true
	BindingName string `url:"bindingName" json:"bindingName"`
}

func (lbrq FrontendBindDeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName can not be empty")
	}
	if lbrq.BindingName == "" {
		return errors.New("validation-error: field BindingName can not be empty")
	}

	return nil
}

// FrontendBindDelete deletes binding from the specified load balancer frontend
func (l LB) FrontendBindDelete(ctx context.Context, req FrontendBindDeleteRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/lb/frontendBindDelete"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
