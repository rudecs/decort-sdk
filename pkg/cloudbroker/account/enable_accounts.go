package account

import (
	"context"
	"errors"
	"net/http"
)

type EnableAccountsRequest struct {
	AccountIDs []uint64 `url:"accountIds"`
}

func (arq EnableAccountsRequest) Validate() error {
	if arq.AccountIDs == nil || len(arq.AccountIDs) == 0 {
		return errors.New("validation-error: field AccountIDs must be set")
	}
	return nil
}

func (a Account) EnableAccounts(ctx context.Context, req EnableAccountsRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/enableAccounts"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
