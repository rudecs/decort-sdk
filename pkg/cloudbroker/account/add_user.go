package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for adding permission to access to account for a user
type AddUserRequest struct {
	// ID of account to add to
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Name of the user to be given rights
	// Required: true
	UserName string `url:"username"`

	// Account permission types:
	//	- 'R' for read only access
	//	- 'RCX' for Write
	//	- 'ARCXDU' for Admin
	// Required: true
	AccessType string `url:"accesstype"`
}

func (arq AddUserRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}
	if arq.UserName == "" {
		return errors.New("validation-error: field UserName can not be empty")
	}
	if arq.AccessType == "" {
		return errors.New("validation-error: field AccessType can not be empty")
	}
	validate := validators.StringInSlice(arq.AccessType, []string{"R", "RCX", "ARCXDU"})
	if !validate {
		return errors.New("validation-error: field AccessType can be only R, RCX or ARCXDU")
	}

	return nil
}

// AddUser gives a user access rights.
func (a Account) AddUser(ctx context.Context, req AddUserRequest) (bool, error) {
	err := req.validate()
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
