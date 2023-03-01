package locations

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of locations
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all locations
func (l Locations) List(ctx context.Context, req ListRequest) (ListLocations, error) {
	url := "/cloudapi/locations/list"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListLocations{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
