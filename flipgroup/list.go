package flipgroup

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (f FlipGroup) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (FlipGroupList, error) {
	url := "/cloudapi/flipgroup/list"
	fgListRaw, err := f.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	fgList := FlipGroupList{}
	if err := json.Unmarshal(fgListRaw, &fgList); err != nil {
		return nil, err
	}

	return fgList, nil
}
