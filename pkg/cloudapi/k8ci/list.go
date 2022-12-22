package k8ci

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list information about images
type ListRequest struct {
	// List disabled items as well
	// Required: false
	IncludeDisabled bool `url:"includeDisabled,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list all k8ci catalog items available to the current user
func (k K8CI) List(ctx context.Context, req ListRequest) (ListK8CI, error) {
	url := "/cloudapi/k8ci/list"

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
