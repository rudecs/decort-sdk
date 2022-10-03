package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type UpdateUserRequest struct {
	AccountID  uint64 `url:"accountId"`
	UserID     string `url:"userId"`
	AccessType string `url:"accesstype"`
}

func (arq UpdateUserRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	if arq.UserID == "" {
		return errors.New("validation-error: field UserID can not be empty")
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

func (a Account) UpdateUser(ctx context.Context, req UpdateUserRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/updateUser"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
