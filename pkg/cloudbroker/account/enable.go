package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for enable account
type EnableRequest struct {
	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Reason to enable
	// Required: true
	Reason string `url:"reason"`
}

func (arq EnableRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Reason == "" {
		return errors.New("field Reason must be set")
	}

	return nil
}

// Enable enables an account
func (a Account) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/enable"

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
