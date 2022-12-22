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
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.Description == "" {
		return errors.New("validation-error: field Description can not be empty")
	}

	return nil
}

// Update updates some of load balancer attributes
func (l LB) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/update"

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
