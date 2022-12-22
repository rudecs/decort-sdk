package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete worker from group
type DeleteWorkerFromGroupRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId"`

	// Compute ID of worker node to delete
	// Required: true
	WorkerID uint64 `url:"workerId"`
}

func (krq DeleteWorkerFromGroupRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkerGroupID must be set")
	}
	if krq.WorkerID == 0 {
		return errors.New("validation-error: field WorkerIDs must be set")
	}

	return nil
}

// DeleteWorkerFromGroup deletes worker compute from workers group in selected kubernetes cluster
func (k K8S) DeleteWorkerFromGroup(ctx context.Context, req DeleteWorkerFromGroupRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/deleteWorkerFromGroup"

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
