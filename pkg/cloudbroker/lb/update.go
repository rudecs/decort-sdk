package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update load balancer
type UpdateRequest struct {
	// ID of the load balancer to update
	// Required: true
	LBID uint64 `url:"lbId"`

	// New description of this load balancer.
	// If omitted, current description is retained
	// Required: true
	Description string `url:"desc"`
}

func (lbrq UpdateRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}
	if lbrq.Description == "" {
		return errors.New("validation-error: field Description must be set")
	}

	return nil
}

// Update updates some of load balancer attributes
func (lb LB) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/update"

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
