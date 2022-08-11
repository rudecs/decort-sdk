package vins

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SearchRequest struct {
	AccountId uint64 `url:"accountId,omitempty"`
	RGID      uint64 `url:"rgId,omitempty"`
	Name      string `url:"name,omitempty"`
	ShowAll   bool   `url:"show_all,omitempty"`
}

func (v Vins) Search(ctx context.Context, req SearchRequest, options ...opts.DecortOpts) (VinsList, error) {
	url := "/vins/search"
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
