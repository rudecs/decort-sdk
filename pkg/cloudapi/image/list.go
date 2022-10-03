package image

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	AccountID uint64 `json:"accountId"`
	Page      uint64 `json:"page"`
	Size      uint64 `json:"size"`
}

func (i Image) List(ctx context.Context, req ListRequest) (ImageList, error) {

	url := "/cloudapi/image/list"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
