package account

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for delete group accounts
type DeleteAccountsRequest struct {
	// IDs of accounts
	// Required: true
	AccountsIDs []uint64 `url:"accountIds"`

	// Reason for deletion
	// Required: true
	Reason string `url:"reason"`

	// Whether to completely destroy accounts or not
	// Required: false
	Permanently bool `url:"permanently,omitempty"`
}

func (arq DeleteAccountsRequest) validate() error {
	if len(arq.AccountsIDs) == 0 {
		return errors.New("validation-error: field AccountIDs must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// DeleteAccounts destroy a group of accounts
func (a Account) DeleteAccounts(ctx context.Context, req DeleteAccountsRequest) (bool, error) {
	err := req.validate()
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
