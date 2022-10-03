package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	K8SID uint64 `url:"k8sId"`
}

func (krq GetRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) Get(ctx context.Context, req GetRequest) (*K8SRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/k8s/get"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	k8sInfo := &K8SRecord{}

	err = json.Unmarshal(res, k8sInfo)
	if err != nil {
		return nil, err
	}

	return k8sInfo, nil
}
