package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includeDeleted"`
	Page           uint64 `url:"page"`
	Size           uint64 `url:"size"`
}

func (v VINS) List(ctx context.Context, req ListRequest) (VINSList, error) {
	url := "/cloudapi/vins/list"

	VINSListRaw, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	VINSList := VINSList{}
	err = json.Unmarshal(VINSListRaw, &VINSList)
	if err != nil {
		return nil, err
	}

	return VINSList, nil

}
