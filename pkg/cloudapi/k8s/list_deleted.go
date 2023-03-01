package k8s

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list deleted kubernetes cluster
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets all deleted kubernetes clusters the user has access to
func (k8s K8S) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListK8SClusters, error) {
	url := "/cloudapi/k8s/listDeleted"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8SClusters{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
