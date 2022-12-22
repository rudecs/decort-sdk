package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable kubernetes cluster
type DisableRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`
}

func (krq DisableRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}

	return nil
}

// Disable disables kubernetes cluster by ID
func (k K8S) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/disable"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
