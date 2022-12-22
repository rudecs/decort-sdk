package flipgroup

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list FLIPGroup available to the current user
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list FLIPGroup managed cluster instances available to the current user
func (f FLIPGroup) List(ctx context.Context, req ListRequest) (ListFLIPGroups, error) {
	url := "/cloudapi/flipgroup/list"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListFLIPGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
