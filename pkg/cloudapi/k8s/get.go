package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get detailed information about kubernetes cluster
type GetRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`
}

func (krq GetRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	return nil
}

// Get gets information about Kubernetes cluster
func (k8s K8S) Get(ctx context.Context, req GetRequest) (*RecordK8S, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/k8s/get"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8S{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
