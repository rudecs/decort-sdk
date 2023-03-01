package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable account
type DisableRequest struct {
	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`

	// Reason to disable
	// Required: true
	Reason string `url:"reason" json:"reason"`
}

func (arq DisableRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// Disable disables an account
func (a Account) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/disable"

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
