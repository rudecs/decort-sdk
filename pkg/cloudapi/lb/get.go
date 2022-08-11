package lb

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
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

func (l LB) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*LoadBalancer, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/lb/get"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	lbRaw, err := l.client.DecortApiCall(ctx, typed.POST, url, req)
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
