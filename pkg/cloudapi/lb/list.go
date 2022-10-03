package lb

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includedeleted"`
	Page           uint64 `url:"page"`
	Size           uint64 `url:"size"`
}

func (l LB) List(ctx context.Context, req ListRequest) (LBList, error) {
	url := "/cloudapi/lb/list"

	lbListRaw, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	lbList := LBList{}
	err = json.Unmarshal(lbListRaw, &lbList)
	if err != nil {
		return nil, err
	}

	return lbList, nil

}
