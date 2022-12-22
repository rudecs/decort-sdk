package grid

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list locations
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list all locations
func (g Grid) List(ctx context.Context, req ListRequest) (ListGrids, error) {
	url := "/cloudbroker/grid/list"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListGrids{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
