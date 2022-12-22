package k8s

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for get configuration of kubernetes cluster
type GetConfigRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`
}

func (krq GetConfigRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	return nil
}

// GetConfig gets configuration data to access Kubernetes cluster
func (k8s K8S) GetConfig(ctx context.Context, req GetConfigRequest) (string, error) {
	err := req.validate()
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
