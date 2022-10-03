package image

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	SepID      uint64 `url:"sepId,omitempty"`
	SharedWith uint64 `url:"sharedWith,omitempty"`
	Page       uint64 `url:"page,omitempty"`
	Size       uint64 `url:"size,omitempty"`
}

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
