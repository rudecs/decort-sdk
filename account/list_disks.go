package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListDisksRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq ListDisksRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListDisks(ctx context.Context, req ListDisksRequest, options ...opts.DecortOpts) (AccountDisksList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/listDisks"
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

	accountDisksList := AccountDisksList{}

	err = json.Unmarshal(res, &accountDisksList)
	if err != nil {
		return nil, err
	}

	return accountDisksList, nil

}
