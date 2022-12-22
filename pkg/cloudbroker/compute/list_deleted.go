package compute

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get deleted computes list
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// ListDeleted gets list all deleted computes
func (c Compute) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListComputes, error) {
	url := "/cloudbroker/compute/listDeleted"

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
