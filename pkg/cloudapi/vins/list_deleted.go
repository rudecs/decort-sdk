package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (v VINS) ListDeleted(ctx context.Context, req ListDeletedRequest) (VINSList, error) {
	url := "/cloudapi/vins/listDeleted"

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
