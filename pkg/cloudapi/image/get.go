package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	ImageID uint64 `url:"imageId"`
	ShowAll bool   `url:"show_all,omitempty"`
}

func (irq GetRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID can not be empty or equal to 0")
	}

	return nil
}

func (i Image) Get(ctx context.Context, req GetRequest) (*ImageExtend, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/image/get"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	imageInfo := &ImageExtend{}

	err = json.Unmarshal(res, imageInfo)
	if err != nil {
		return nil, err
	}

	return imageInfo, nil
}
