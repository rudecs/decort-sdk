package account

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for restore a deleted account
type RestoreRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Reason to restore
	// Required: true
	Reason string `url:"reason"`
}

func (arq RestoreRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// Restore restores a deleted account
func (a Account) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/restore"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
