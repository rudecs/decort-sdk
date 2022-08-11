package image

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	AccountId uint64 `json:"accountId"`
	Page      uint64 `json:"page"`
	Size      uint64 `json:"size"`
}

func (i Image) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (ImageList, error) {

	url := "/image/list"
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

	imageList := ImageList{}

	err = json.Unmarshal(res, &imageList)
	if err != nil {
		return nil, err
	}

	return imageList, nil
}
