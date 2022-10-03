package k8s

import (
	"context"
	"errors"
	"net/http"
)

type GetConfigRequest struct {
	K8SID uint64 `url:"k8sId"`
}

func (krq GetConfigRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) GetConfig(ctx context.Context, req GetConfigRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/k8s/getConfig"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
