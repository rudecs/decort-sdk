package sizes

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for list the available flavors
type ListRequest struct {
	// ID of the cloudspace
	// Required: false
	CloudspaceID uint64 `url:"cloudspaceId,omitempty" json:"cloudspaceId,omitempty"`

	// Location code for the sizes
	// Required: false
	Location string `url:"location,omitempty" json:"location,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list the available flavors, filtering can be based on the user which is doing the request
func (s Sizes) List(ctx context.Context, req ListRequest) (ListSizes, error) {
	url := "/cloudapi/sizes/list"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSizes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
