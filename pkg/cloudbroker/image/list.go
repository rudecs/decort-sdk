package image

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list available images
type ListRequest struct {
	// Filter images by storage endpoint provider ID
	// Required: false
	SepID uint64 `url:"sepId,omitempty"`

	// Filter images by account ID availability
	// Required: false
	SharedWith uint64 `url:"sharedWith,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list of information about images
func (i Image) List(ctx context.Context, req ListRequest) (ListImages, error) {
	url := "/cloudbroker/image/list"

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
