package rg

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of resource groups
type ListRequest struct {
	// Included deleted resource groups
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list of all resource groups the user has access to
func (r RG) List(ctx context.Context, req ListRequest) (ListResourceGroups, error) {
	url := "/cloudapi/rg/list"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListResourceGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
