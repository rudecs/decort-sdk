package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restart worker node
type WorkerRestartRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId"`

	// Compute ID of worker node to restart
	// Required: true
	WorkerID uint64 `url:"workerId"`
}

func (krq WorkerRestartRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID can not be empty or equal to 0")
	}
	if krq.WorkerID == 0 {
		return errors.New("validation-error: field WorkerID can not be empty or equal to 0")
	}

	return nil
}

// WorkerRestart soft restart (reboot OS) worker node of the Kubernetes cluster
func (k8s K8S) WorkerRestart(ctx context.Context, req WorkerRestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/workerRestart"

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
