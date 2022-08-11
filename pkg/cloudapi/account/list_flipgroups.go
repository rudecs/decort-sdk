package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListFlipGroupsRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq ListFlipGroupsRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListFlipGroups(ctx context.Context, req ListFlipGroupsRequest, options ...opts.DecortOpts) (AccountFlipGroupsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/listFlipGroups"
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

	accountFlipGroupsList := AccountFlipGroupsList{}

	err = json.Unmarshal(res, &accountFlipGroupsList)
	if err != nil {
		return nil, err
	}

	return accountFlipGroupsList, nil

}
