package account

import (
	"context"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetConsumtionRequest struct {
	AccountId uint64 `url:"accountId"`
	Start     uint64 `url:"start"`
	End       uint64 `url:"end"`
}

func (arq GetConsumtionRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	if arq.Start == 0 {
		return errors.New("validation-error: field Start can not be empty or equal to 0")
	}

	if arq.End == 0 {
		return errors.New("validation-error: field End can not be empty or equal to 0")
	}

	return nil
}

func (a Account) GetConsumtion(ctx context.Context, req GetConsumtionRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/account/getConsumtion"
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
		return "", err
	}

	return string(res), nil

}

func (a Account) GetConsumtionGet(ctx context.Context, req GetConsumtionRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/account/getConsumtion"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := a.client.DecortApiCall(ctx, typed.GET, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}
