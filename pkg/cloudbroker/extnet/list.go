package extnet

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list external network
type ListRequest struct {
	// Filter by account ID
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all available external networks
func (e ExtNet) List(ctx context.Context, req ListRequest) (ListExtNet, error) {
	url := "/cloudbroker/extnet/list"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListExtNet{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
