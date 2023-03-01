package lb

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for create load balancer
type CreateRequest struct {
	// ID of the resource group where this load balancer instance will be located
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// Name of the load balancer.
	// Must be unique among all load balancers in this Resource Group
	// Required: true
	Name string `url:"name" json:"name"`

	// External network to connect this load balancer to
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId"`

	// Internal network (VINS) to connect this load balancer to
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Start now Load balancer
	// Required: false
	Start bool `url:"start" json:"start"`

	// Text description of this load balancer
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

func (lbrq CreateRequest) validate() error {
	if lbrq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if lbrq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if lbrq.ExtNetID == 0 {
		return errors.New("validation-error: field ExtNetID can not be empty or equal to 0")
	}
	if lbrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

// Create method will create a new load balancer instance
func (l LB) Create(ctx context.Context, req CreateRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/lb/create"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
