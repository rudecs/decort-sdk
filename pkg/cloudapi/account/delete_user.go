package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for revoke access to account
type DeleteUserRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`

	// ID or emailaddress of the user to remove
	// Required: true
	UserID string `url:"userId" json:"userId"`

	// Recursively revoke access rights from owned cloudspaces and vmachines
	// Required: false
	RecursiveDelete bool `url:"recursivedelete,omitempty" json:"recursivedelete,omitempty"`
}

func (arq DeleteUserRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}
	if arq.UserID == "" {
		return errors.New("validation-error: field UserID can not be empty")
	}

	return nil
}

// DeleteUser revokes user access from the account
func (a Account) DeleteUser(ctx context.Context, req DeleteUserRequest) (bool, error) {
	err := req.validate()
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
