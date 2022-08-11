package account

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteUserRequest struct {
	AccountId       uint64 `url:"accountId"`
	UserId          string `url:"userId"`
	RecursiveDelete bool   `url:"recursivedelete,omitempty"`
}

func (arq DeleteUserRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	if arq.UserId == "" {
		return errors.New("validation-error: field UserId can not be empty")
	}

	return nil
}

func (a Account) DeleteUser(ctx context.Context, req DeleteUserRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/account/deleteUser"
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
