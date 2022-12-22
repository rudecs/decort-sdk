package lb

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for frontend bind
type FrontendBindRequest struct {
	// ID of the load balancer instance to FrontendBind
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

func (lbrq FrontendBindRequest) validate() error {
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

// FrontendBind bind frontend from specified load balancer instance
func (l LB) FrontendBind(ctx context.Context, req FrontendBindRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/lb/frontendBind"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
