package vins

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (v Vins) ListDeleted(ctx context.Context, req ListDeletedRequest, options ...opts.DecortOpts) (VinsList, error) {
	url := "/vins/listDeleted"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	vinsListRaw, err := v.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	vinsList := VinsList{}
	err = json.Unmarshal(vinsListRaw, &vinsList)
	if err != nil {
		return nil, err
	}

	return vinsList, nil

}
