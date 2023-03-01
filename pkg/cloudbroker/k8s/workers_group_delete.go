package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete workers group
type WorkersGroupDeleteRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// Worker group ID
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId" json:"workersGroupId"`
}

func (krq WorkersGroupDeleteRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID must be set")
	}

	return nil
}

// WorkersGroupDelete deletes workers group from kubernetes cluster
func (k K8S) WorkersGroupDelete(ctx context.Context, req WorkersGroupDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/workersGroupDelete"

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
