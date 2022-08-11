package account

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (a Account) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (AccountCloudApiList, error) {
	url := "/account/list"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := a.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	accountList := AccountCloudApiList{}

	err = json.Unmarshal(res, &accountList)
	if err != nil {
		return nil, err
	}

	return accountList, nil

}
