package compute

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list available computes
type ListRequest struct {
	// Include deleted computes
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty" json:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of the available computes.
// Filtering based on status is possible
func (c Compute) List(ctx context.Context, req ListRequest) (ListComputes, error) {
	url := "/cloudapi/compute/list"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
