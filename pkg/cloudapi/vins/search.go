package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	AccountID uint64 `url:"accountId,omitempty"`
	RGID      uint64 `url:"rgId,omitempty"`
	Name      string `url:"name,omitempty"`
	ShowAll   bool   `url:"show_all,omitempty"`
}

func (v VINS) Search(ctx context.Context, req SearchRequest) (VINSList, error) {
	url := "/cloudapi/vins/search"

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
