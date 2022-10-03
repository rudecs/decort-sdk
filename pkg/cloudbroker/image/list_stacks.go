package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListStacksRequest struct {
	ImageID uint64 `url:"imageId"`
	Page    uint64 `url:"page,omitempty"`
	Size    uint64 `url:"size,omitempty"`
}

func (irq ListStacksRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (i Image) ListStacks(ctx context.Context, req ListStacksRequest) (ListStacks, error) {
	err := req.Validate()
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
