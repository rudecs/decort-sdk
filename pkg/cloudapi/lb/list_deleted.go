package lb

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of deleted load balancers
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: true
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list of deleted load balancers
func (l LB) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListLB, error) {
	url := "/cloudapi/lb/listDeleted"

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
