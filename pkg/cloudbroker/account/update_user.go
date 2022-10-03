package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateUserRequest struct {
	AccountID  uint64 `url:"accountId"`
	UserID     string `url:"userId"`
	AccessType string `url:"accesstype"`
}

func (arq UpdateUserRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.UserID == "" {
		return errors.New("validation-error: field UserID must be set")
	}
	if arq.AccessType == "" {
		return errors.New("validation-error: field AccessType must be set")
	}

	return nil
}

func (a Account) UpdateUser(ctx context.Context, req UpdateUserRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/updateUser"

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
