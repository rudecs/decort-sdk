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
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq GetRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

// Get gets detailed information about load balancer
func (l LB) Get(ctx context.Context, req GetRequest) (*RecordLB, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/lb/get"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
