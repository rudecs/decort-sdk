package lb

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	LBID uint64 `url:"lbId"`
}

func (lbrq GetRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

func (l LB) Get(ctx context.Context, req GetRequest) (*LoadBalancer, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/lb/get"

	lbRaw, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	lb := &LoadBalancer{}
	err = json.Unmarshal(lbRaw, lb)
	if err != nil {
		return nil, err
	}

	return lb, nil

}
