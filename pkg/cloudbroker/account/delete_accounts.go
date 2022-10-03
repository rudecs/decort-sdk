package account

import (
	"context"
	"errors"
	"net/http"
)

type DeleteAccountsRequest struct {
	AccountsIDs []uint64 `url:"accountIds"`
	Reason      string   `url:"reason"`
	Permanently bool     `url:"permanently,omitempty"`
}

func (arq DeleteAccountsRequest) Validate() error {
	if arq.AccountsIDs == nil || len(arq.AccountsIDs) == 0 {
		return errors.New("validation-error: field AccountIDs must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}
	return nil
}

func (a Account) DeleteAccounts(ctx context.Context, req DeleteAccountsRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/deleteAccounts"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
