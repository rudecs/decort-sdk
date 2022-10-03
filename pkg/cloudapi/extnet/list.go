package extnet

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	AccountID uint64 `url:"accountId"`
	Page      uint64 `url:"page"`
	Size      uint64 `url:"size"`
}

func (e ExtNet) List(ctx context.Context, req ListRequest) (ExtNetList, error) {
	url := "/cloudapi/extnet/list"

	extnetListRaw, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	extnetList := ExtNetList{}
	err = json.Unmarshal(extnetListRaw, &extnetList)
	if err != nil {
		return nil, err
	}

	return extnetList, nil

}
