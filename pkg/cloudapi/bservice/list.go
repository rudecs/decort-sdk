package bservice

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	AccountID uint64 `url:"accountId,omitempty"`
	RGID      uint64 `url:"rgId,omitempty"`
	Page      uint64 `url:"page,omitempty"`
	Size      uint64 `url:"size,omitempty"`
}

func (b BService) List(ctx context.Context, req ListRequest) (BasicServiceList, error) {
	url := "/cloudapi/bservice/list"
	bsListRaw, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	bsList := BasicServiceList{}
	if err := json.Unmarshal(bsListRaw, &bsList); err != nil {
		return nil, err
	}

	return bsList, nil
}

func (b BService) ListDeleted(ctx context.Context, req ListRequest) (BasicServiceList, error) {
	url := "/cloudapi/bservice/listDeleted"
	bsListRaw, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	bsList := BasicServiceList{}
	if err := json.Unmarshal(bsListRaw, &bsList); err != nil {
		return nil, err
	}

	return bsList, nil
}
