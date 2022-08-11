package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRGRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq ListRGRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListRG(ctx context.Context, req ListRGRequest, options ...opts.DecortOpts) (AccountRGList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/listRG"
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

	accountRGList := AccountRGList{}

	err = json.Unmarshal(res, &accountRGList)
	if err != nil {
		return nil, err
	}

	return accountRGList, nil

}
