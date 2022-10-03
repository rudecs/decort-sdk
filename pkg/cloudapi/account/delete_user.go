package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteUserRequest struct {
	AccountID       uint64 `url:"accountId"`
	UserID          string `url:"userId"`
	RecursiveDelete bool   `url:"recursivedelete,omitempty"`
}

func (arq DeleteUserRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	if arq.UserID == "" {
		return errors.New("validation-error: field UserID can not be empty")
	}

	return nil
}

func (a Account) DeleteUser(ctx context.Context, req DeleteUserRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/deleteUser"

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
