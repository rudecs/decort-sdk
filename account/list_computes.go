package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListComputesRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq ListComputesRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListComputes(ctx context.Context, req ListComputesRequest, options ...opts.DecortOpts) (AccountComputesList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/listComputes"
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

	accountComputesList := AccountComputesList{}

	err = json.Unmarshal(res, &accountComputesList)
	if err != nil {
		return nil, err
	}

	return accountComputesList, nil

}
