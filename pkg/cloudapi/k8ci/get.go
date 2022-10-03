package k8ci

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	K8CIID uint64 `url:"k8ciId"`
}

func (krq GetRequest) Validate() error {
	if krq.K8CIID == 0 {
		return errors.New("field K8CIID can not be empty or equal to 0")
	}

	return nil
}

func (k K8CI) Get(ctx context.Context, req GetRequest) (*K8CIRecord, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/k8ci/get"
	k8ciRaw, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	k8ci := &K8CIRecord{}
	if err := json.Unmarshal(k8ciRaw, k8ci); err != nil {
		return nil, err
	}

	return k8ci, nil
}
