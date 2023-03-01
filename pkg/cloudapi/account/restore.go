package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restore a deleted account
type RestoreRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq RestoreRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// Restore restores a deleted account
func (a Account) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/restore"

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
