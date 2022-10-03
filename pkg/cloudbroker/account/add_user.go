package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type AddUserRequest struct {
	AccountID  uint64 `url:"accountId"`
	UserName   string `url:"username"`
	AccessType string `url:"accesstype"`
}

func (arq AddUserRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	if arq.UserName == "" {
		return errors.New("validation-error: field UserName can not be empty")
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

func (a Account) AddUser(ctx context.Context, req AddUserRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/addUser"

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