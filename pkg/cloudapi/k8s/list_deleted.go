package k8s

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (k8s K8S) ListDeleted(ctx context.Context, req ListDeletedRequest) (K8SList, error) {

	url := "/cloudapi/k8s/listDeleted"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	k8sList := K8SList{}

	err = json.Unmarshal(res, &k8sList)
	if err != nil {
		return nil, err
	}

	return k8sList, nil
}
