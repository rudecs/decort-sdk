package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about worker group
type WorkersGroupGetByNameRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// Worker group name
	// Required: true
	GroupName string `url:"groupName"`
}

func (krq WorkersGroupGetByNameRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.GroupName == "" {
		return errors.New("validation-error: field WorkersGroupID must be set")
	}

	return nil
}

// WorkersGroupGetByName gets worker group metadata by name
func (k K8S) WorkersGroupGetByName(ctx context.Context, req WorkersGroupGetByNameRequest) (*RecordK8SGroup, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/k8s/workersGroupGetByName"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8SGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
