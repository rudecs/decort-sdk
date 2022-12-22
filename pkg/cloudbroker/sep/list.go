package sep

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of SEPs
type ListRequest struct {
	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`
}

// List gets list of SEPs
func (s SEP) List(ctx context.Context, req ListRequest) (ListSEP, error) {
	url := "/cloudbroker/sep/list"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSEP{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
