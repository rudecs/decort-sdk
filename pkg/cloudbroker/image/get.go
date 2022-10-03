package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	ImageID uint64 `url:"imageId"`
}

func (irq GetRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (i Image) Get(ctx context.Context, req GetRequest) (*ImageRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/image/get"

	result := ImageRecord{}

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
