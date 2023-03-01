package image

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list available images
type ListRequest struct {
	// Optional account ID to include account images
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list available images, optionally filtering by account ID
func (i Image) List(ctx context.Context, req ListRequest) (ListImages, error) {
	url := "/cloudapi/image/list"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListImages{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
