package lb

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get detailed information about load balancer
type GetRequest struct {
	// ID of the load balancer to get details for
	// Required: true
	LBID uint64 `url:"lbId"`
}

func (lbrq GetRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	return nil
}

// Get gets detailed information about load balancer
func (lb LB) Get(ctx context.Context, req GetRequest) (*RecordLB, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/lb/get"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordLB{}
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil

}
