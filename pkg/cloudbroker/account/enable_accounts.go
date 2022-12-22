package account

import (
	"context"
	"errors"
	"net/http"
)

// Request for enable group accounts
type EnableAccountsRequest struct {
	// IDs od accounts
	// Required: true
	AccountIDs []uint64 `url:"accountIds"`
}

func (arq EnableAccountsRequest) validate() error {
	if len(arq.AccountIDs) == 0 {
		return errors.New("validation-error: field AccountIDs must be set")
	}

	return nil
}

// EnableAccounts enables accounts
func (a Account) EnableAccounts(ctx context.Context, req EnableAccountsRequest) (bool, error) {
	err := req.validate()
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
