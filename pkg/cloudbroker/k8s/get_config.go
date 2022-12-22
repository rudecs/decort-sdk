package k8s

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for get configuration of kubernetes cluster
type GetConfigRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`
}

func (krq GetConfigRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}

	return nil
}

// GetConfig gets configuration data to access kubernetes cluster
func (k K8S) GetConfig(ctx context.Context, req GetConfigRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/k8s/getConfig"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
