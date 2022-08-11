package bservice

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq GetRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*BasicService, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/bservice/get"
	bsRaw, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	bs := &BasicService{}
	if err := json.Unmarshal(bsRaw, bs); err != nil {
		return nil, err
	}

	return bs, nil
}
