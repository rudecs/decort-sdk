package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq GetRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*AccountWithResources, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/get"
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

	account := &AccountWithResources{}

	err = json.Unmarshal(res, &account)
	if err != nil {
		return nil, err
	}

	return account, nil

}
