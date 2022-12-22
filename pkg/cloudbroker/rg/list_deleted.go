package rg

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list deleted resource groups
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// ListDeleted gets list all deleted resource groups the user has access to
func (r RG) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListRG, error) {
	url := "/cloudbroker/rg/listDeleted"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListRG{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
