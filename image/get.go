package image

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	ImageId uint64 `url:"imageId"`
	ShowAll bool   `url:"show_all,omitempty"`
}

func (irq GetRequest) Validate() error {
	if irq.ImageId == 0 {
		return errors.New("validation-error: field ImageId can not be empty or equal to 0")
	}

	return nil
}

func (i Image) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*ImageExtend, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/image/get"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := i.client.DecortApiCall(ctx, typed.POST, url, req)
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
