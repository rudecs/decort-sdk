package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create load balancer
type CreateRequest struct {
	// ID of the resource group where this load balancer instance will be located
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// Name of the load balancer.
	// Must be unique among all load balancers in this resource group
	// Required: true
	Name string `url:"name" json:"name"`

	// OS image ID to create load balancer from
	// Required: false
	ImageID uint64 `url:"imageId,omitempty" json:"imageId,omitempty"`

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
		return errors.New("validation-error: field RGID must be set")
	}
	if lbrq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if lbrq.ExtNetID == 0 {
		return errors.New("validation-error: field ExtNetID must be set")
	}
	if lbrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// Create method will create a new load balancer instance
func (lb LB) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/lb/create"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
