package bservice

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list/deleted list BasicService instances
type ListRequest struct {
	// ID of the account to query for BasicService instances
	// Required: false
	AccountID uint64 `url:"accountId,omitempty"`

	// ID of the resource group to query for BasicService instances
	// Required: false
	RGID uint64 `url:"rgId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list BasicService instances associated with the specified Resource Group
func (b BService) List(ctx context.Context, req ListRequest) (ListBasicServices, error) {
	url := "/cloudapi/bservice/list"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListBasicServices{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// ListDeleted gets list deleted BasicService instances associated with the specified Resource Group
func (b BService) ListDeleted(ctx context.Context, req ListRequest) (ListBasicServices, error) {
	url := "/cloudapi/bservice/listDeleted"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListBasicServices{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
