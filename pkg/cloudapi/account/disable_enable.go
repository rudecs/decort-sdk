package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for change status of account
type DisabelEnableRequest struct {
	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq DisabelEnableRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// Disable disables an account
func (a Account) Disable(ctx context.Context, req DisabelEnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/disable"

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

// Enable enables an account
func (a Account) Enable(ctx context.Context, req DisabelEnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/enable"

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
