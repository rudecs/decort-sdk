package account

import (
	"context"
	"errors"
	"net/http"
)

type DisableAccountsRequest struct {
	AccountIDs []uint64 `url:"accountIds,omitempty"`
}

func (arq DisableAccountsRequest) Validate() error {
	if arq.AccountIDs == nil || len(arq.AccountIDs) == 0 {
		return errors.New("validation-error: field AccountIDs must be set")
	}
	return nil
}

func (a Account) DisableAccounts(ctx context.Context, req DisableAccountsRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/disableAccounts"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
