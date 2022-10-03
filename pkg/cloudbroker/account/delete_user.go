package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteUserRequest struct {
	AccountID       uint64 `url:"accountId"`
	UserName        string `url:"username"`
	RecursiveDelete bool   `url:"recursivedelete,omitempty"`
}

func (arq DeleteUserRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.UserName == "" {
		return errors.New("validation-error: field UserName must be set")
	}

	return nil
}

func (a Account) DeleteUser(ctx context.Context, req DeleteUserRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/deleteUser"

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
