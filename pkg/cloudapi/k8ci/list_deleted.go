package k8ci

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list information about deleted images
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// ListDeleted gets list all deleted k8ci catalog items available to the current user
func (k K8CI) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListK8CI, error) {
	url := "/cloudapi/k8ci/listDeleted"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8CI{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
