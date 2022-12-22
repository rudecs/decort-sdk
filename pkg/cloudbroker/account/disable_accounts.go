package account

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for disable group accounts
type DisableAccountsRequest struct {
	// IDs of accounts
	// Required: true
	AccountIDs []uint64 `url:"accountIds,omitempty"`
}

func (arq DisableAccountsRequest) validate() error {
	if len(arq.AccountIDs) == 0 {
		return errors.New("validation-error: field AccountIDs must be set")
	}

	return nil
}

// DisableAccounts disables accounts
func (a Account) DisableAccounts(ctx context.Context, req DisableAccountsRequest) (bool, error) {
	err := req.validate()
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
