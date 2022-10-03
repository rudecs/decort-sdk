package k8ci

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (k K8CI) ListDeleted(ctx context.Context, req ListDeletedRequest) (K8CIList, error) {
	url := "/cloudapi/k8ci/listDeleted"
	k8ciListRaw, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	k8ciList := K8CIList{}
	if err := json.Unmarshal(k8ciListRaw, &k8ciList); err != nil {
		return nil, err
	}

	return k8ciList, nil
}
