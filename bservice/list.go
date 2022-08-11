package bservice

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	AccountID uint64 `url:"accountId,omitempty"`
	RGID      uint64 `url:"rgId,omitempty"`
	Page      uint64 `url:"page,omitempty"`
	Size      uint64 `url:"size,omitempty"`
}

func (b BService) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (BasicServiceList, error) {
	url := "/cloudapi/bservice/list"
	bsListRaw, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	bsList := BasicServiceList{}
	if err := json.Unmarshal(bsListRaw, &bsList); err != nil {
		return nil, err
	}

	return bsList, nil
}

func (b BService) ListDeleted(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (BasicServiceList, error) {
	url := "/cloudapi/bservice/listDeleted"
	bsListRaw, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	bsList := BasicServiceList{}
	if err := json.Unmarshal(bsListRaw, &bsList); err != nil {
		return nil, err
	}

	return bsList, nil
}
