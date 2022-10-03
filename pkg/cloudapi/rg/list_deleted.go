package rg

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (r RG) ListDeleted(ctx context.Context, req ListDeletedRequest) (ResourceGroupList, error) {
	url := "/cloudapi/rg/listDeleted"
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
