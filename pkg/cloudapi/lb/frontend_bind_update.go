package lb

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for update binding
type FrontendBindUpdateRequest struct {
	// ID of the load balancer instance to FrontendBindUpdate
	// Required: true
	LBID uint64 `url:"lbId"`

	// Name of the frontend to update
	// Required: true
	FrontendName string `url:"frontendName"`

	// Name of the binding to update
	// Required: true
	BindingName string `url:"bindingName"`

	// If specified must be within the IP range of either Ext Net or ViNS,
	// where this load balancer is connected - new IP address to use for this binding.
	// If omitted, current IP address is retained
	// Required: false
	BindingAddress string `url:"bindingAddress,omitempty"`

	// New port number to use for this binding.
	// If omitted, current port number is retained
	// Required: false
	BindingPort uint64 `url:"bindingPort,omitempty"`
}

func (lbrq FrontendBindUpdateRequest) validate() error {
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

// FrontendBindUpdate updates binding for the specified load balancer frontend
func (l LB) FrontendBindUpdate(ctx context.Context, req FrontendBindUpdateRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/lb/frontendBindingUpdate"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
