package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of VINSes
type ListRequest struct {
	// Include deleted
	// Required: false
	IncludeDeleted bool `url:"includeDeleted,omitempty" json:"includeDeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of VINSes available for current user
func (v VINS) List(ctx context.Context, req ListRequest) (ListVINS, error) {
	url := "/cloudapi/vins/list"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListVINS{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
