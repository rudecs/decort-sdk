package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type WorkersGroupGetByNameRequest struct {
	K8SID     uint64 `url:"k8sId"`
	GroupName string `url:"groupName "`
}

func (krq WorkersGroupGetByNameRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	if krq.GroupName == "" {
		return errors.New("validation-error: field WorkersGroupID can not be empty")
	}

	return nil
}

func (k8s K8S) WorkersGroupGetByName(ctx context.Context, req WorkersGroupGetByNameRequest) (*K8SGroup, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/k8s/workersGroupGetByName"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	group := &K8SGroup{}

	err = json.Unmarshal(res, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}
