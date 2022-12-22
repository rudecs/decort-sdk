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
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID must be set")
	}
	if krq.WorkerID == 0 {
		return errors.New("validation-error: field WorkerID must be set")
	}

	return nil
}

// WorkerRestart soft restart (reboot OS) worker node of the kubernetes cluster
func (k8s K8S) WorkerRestart(ctx context.Context, req WorkerRestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/workerRestart"

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
