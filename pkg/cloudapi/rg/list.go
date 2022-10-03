package rg

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includedeleted,omitempty"`
	Page           uint64 `url:"page,omitempty"`
	Size           uint64 `url:"size,omitempty"`
}

func (r RG) List(ctx context.Context, req ListRequest) (ResourceGroupList, error) {
	url := "/cloudapi/rg/list"
	rgListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	rgList := ResourceGroupList{}
	if err := json.Unmarshal(rgListRaw, &rgList); err != nil {
		return nil, err
	}

	return rgList, nil
}
