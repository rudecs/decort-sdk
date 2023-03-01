package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete load balancer
type DeleteRequest struct {
	// ID of the load balancer instance to delete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`

	// Set to true to delete load balancer immediately bypassing recycle bin
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

func (lbrq DeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Delete deletes specified load balancer
func (lb LB) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/delete"

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
