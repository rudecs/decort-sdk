package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update binding
type FrontendBindUpdateRequest struct {
	// ID of the load balancer instance to FrontendBindUpdate
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`

	// Name of the frontend to update
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName"`

	// Name of the binding to update
	// Required: true
	BindingName string `url:"bindingName" json:"bindingName"`

	// If specified must be within the IP range of either Ext Net or ViNS,
	// where this load balancer is connected - new IP address to use for this binding.
	// If omitted, current IP address is retained
	// Required: false
	BindingAddress string `url:"bindingAddress,omitempty" json:"bindingAddress,omitempty"`

	// New port number to use for this binding.
	// If omitted, current port number is retained
	// Required: false
	BindingPort uint64 `url:"bindingPort,omitempty" json:"bindingPort,omitempty"`
}

func (lbrq FrontendBindUpdateRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}
	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName must be set")
	}
	if lbrq.BindingName == "" {
		return errors.New("validation-error: field BindingName must be set")
	}
	if lbrq.BindingAddress == "" {
		return errors.New("validation-error: field BindingAddress must be set")
	}
	if lbrq.BindingPort == 0 {
		return errors.New("validation-error: field BindingPort must be set")
	}

	return nil
}

// FrontendBindUpdate updates binding for the specified load balancer frontend
func (lb LB) FrontendBindUpdate(ctx context.Context, req FrontendBindUpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/frontendBindingUpdate"

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
