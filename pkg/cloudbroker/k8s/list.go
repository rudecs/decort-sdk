package k8s

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list information K8S
type ListRequest struct {
	// Include deleted clusters in result
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list all kubernetes clusters
func (k K8S) List(ctx context.Context, req ListRequest) (ListK8S, error) {

	url := "/cloudbroker/k8s/list"

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
