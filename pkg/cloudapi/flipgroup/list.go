package flipgroup

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (f FlipGroup) List(ctx context.Context, req ListRequest) (FlipGroupList, error) {
	url := "/cloudapi/flipgroup/list"
	fgListRaw, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	fgList := FlipGroupList{}
	if err := json.Unmarshal(fgListRaw, &fgList); err != nil {
		return nil, err
	}

	return fgList, nil
}
