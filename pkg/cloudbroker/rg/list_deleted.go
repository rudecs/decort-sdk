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

func (r RG) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListDeleted, error) {
	url := "/cloudbroker/rg/listDeleted"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	listDeleted := ListDeleted{}

	err = json.Unmarshal(res, &listDeleted)
	if err != nil {
		return nil, err
	}

	return listDeleted, nil
}
