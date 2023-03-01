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
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// Worker group name
	// Required: true
	GroupName string `url:"groupName" json:"groupName"`
}

func (krq WorkersGroupGetByNameRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.GroupName == "" {
		return errors.New("validation-error: field WorkersGroupID can not be empty")
	}

	return nil
}

// WorkersGroupGetByName gets worker group metadata by name
func (k8s K8S) WorkersGroupGetByName(ctx context.Context, req WorkersGroupGetByNameRequest) (*RecordK8SGroups, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/k8s/workersGroupGetByName"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8SGroups{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
