package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of deleted VINSes
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list of deleted VINSes available for current user
func (v VINS) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListVINS, error) {
	url := "/cloudapi/vins/listDeleted"

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
