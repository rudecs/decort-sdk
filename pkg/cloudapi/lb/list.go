package lb

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of load balancers
type ListRequest struct {
	// Included deleted load balancers
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list all load balancers
func (l LB) List(ctx context.Context, req ListRequest) (ListLB, error) {
	url := "/cloudapi/lb/list"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListLB{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
