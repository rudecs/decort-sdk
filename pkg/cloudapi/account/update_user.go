package account

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type UpdateUserRequest struct {
	AccountId  uint64 `url:"accountId"`
	UserId     string `url:"userId"`
	AccessType string `url:"accesstype"`
}

func (arq UpdateUserRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	if arq.UserId == "" {
		return errors.New("validation-error: field UserId can not be empty")
	}

	if arq.AccessType == "" {
		return errors.New("validation-error: field AccessType can not be empty")
	}

	isValid := validators.StringInSlice(arq.AccessType, []string{"R", "RCX", "ARCXDU"})
	if !isValid {
		return errors.New("validation-error: field AccessType can be only R, RCX or ARCXDU")
	}

	return nil
}

func (a Account) UpdateUser(ctx context.Context, req UpdateUserRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/account/updateUser"
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
