package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list stack
type ListStacksRequest struct {
	// Image ID
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

func (irq ListStacksRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// ListStacks gets list stack by image ID
func (i Image) ListStacks(ctx context.Context, req ListStacksRequest) (ListStacks, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/image/listStacks"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListStacks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
