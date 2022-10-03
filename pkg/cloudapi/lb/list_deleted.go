package lb

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (l LB) ListDeleted(ctx context.Context, req ListDeletedRequest) (LBList, error) {
	url := "/cloudapi/lb/listDeleted"

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
