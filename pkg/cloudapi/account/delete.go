package account

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteRequest struct {
	AccountId   uint64 `url:"accountId"`
	Permanently bool   `url:"permanently,omitempty"`
}

func (arq DeleteRequest) Validate() error {

	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId must be set")
	}
	return nil
}

func (a Account) Delete(ctx context.Context, req DeleteRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/account/delete"
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
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
