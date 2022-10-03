package account

import (
	"context"
	"errors"
	"net/http"
)

type DeleteRequest struct {
	AccountID   uint64 `url:"accountId"`
	Reason      string `url:"reason"`
	Permanently bool   `url:"permanently,omitempty"`
}

func (arq DeleteRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}
	return nil
}

func (a Account) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/delete"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
