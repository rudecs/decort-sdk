package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update user access rights
type UpdateUserRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Userid/Email for registered users or emailaddress for unregistered users
	// Required: true
	UserID string `url:"userId"`

	// Account permission types:
	//	- 'R' for read only access
	//	- 'RCX' for Write
	//	- 'ARCXDU' for Admin
	// Required: true
	AccessType string `url:"accesstype"`
}

func (arq UpdateUserRequest) validate() error {
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

// UpdateUser updates user access rights
func (a Account) UpdateUser(ctx context.Context, req UpdateUserRequest) (bool, error) {
	err := req.validate()
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
