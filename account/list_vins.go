package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListVinsRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq ListVinsRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListVins(ctx context.Context, req ListVinsRequest, options ...opts.DecortOpts) (AccountVinsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/listVins"
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

	accountVinsList := AccountVinsList{}

	err = json.Unmarshal(res, &accountVinsList)
	if err != nil {
		return nil, err
	}

	return accountVinsList, nil

}
