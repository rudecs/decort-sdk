package account

import (
	"context"
	"errors"
	"net/http"
)

type RestoreRequest struct {
	AccountID uint64 `url:"accountId"`
	Reason    string `url:"reason"`
}

func (arq RestoreRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

func (a Account) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.Validate()
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