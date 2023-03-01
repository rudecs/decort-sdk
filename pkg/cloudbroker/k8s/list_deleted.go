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

// ListDeleted gets all deleted kubernetes clusters
func (k K8S) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListK8S, error) {

	url := "/cloudbroker/k8s/listDeleted"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8S{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
