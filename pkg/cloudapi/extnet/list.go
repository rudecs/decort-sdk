package extnet

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	AccountId uint64 `url:"accountId"`
	Page      uint64 `url:"page"`
	Size      uint64 `url:"size"`
}

func (e Extnet) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (ExtnetList, error) {
	url := "/extnet/list"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	extnetListRaw, err := e.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	extnetList := ExtnetList{}
	err = json.Unmarshal(extnetListRaw, &extnetList)
	if err != nil {
		return nil, err
	}

	return extnetList, nil

}
