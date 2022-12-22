package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID must be set")
	}

	return nil
}

// WorkerAdd adds worker nodes to a kubernetes cluster
func (k K8S) WorkerAdd(ctx context.Context, req WorkerAddRequest) ([]uint64, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/k8s/workerAdd"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	result := make([]uint64, 0)

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
