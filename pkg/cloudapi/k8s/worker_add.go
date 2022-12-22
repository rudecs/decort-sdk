package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for add worker to a kubernetes cluster
type WorkerAddRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId"`

	// How many worker nodes to add
	// Required: true
	Num uint64 `url:"num"`
}

func (krq WorkerAddRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID can not be empty or equal to 0")
	}
	if krq.Num == 0 {
		return errors.New("validation-error: field Num can not be empty or equal to 0")
	}

	return nil
}

// WorkerAdd add worker nodes to a Kubernetes cluster
func (k8s K8S) WorkerAdd(ctx context.Context, req WorkerAddRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/workerAdd"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
