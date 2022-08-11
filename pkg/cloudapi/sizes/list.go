package sizes

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	CloudspaceID uint64 `url:"cloudspaceId,omitempty"`
	Location     string `url:"location,omitempty"`
	Page         uint64 `url:"page,omitempty"`
	Size         uint64 `url:"size,omitempty"`
}

func (s Sizes) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (SizesList, error) {
	url := "/cloudapi/sizes/list"
	sizesListRaw, err := s.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	sizesList := SizesList{}
	if err := json.Unmarshal(sizesListRaw, &sizesList); err != nil {
		return nil, err
	}

	return sizesList, nil
}
