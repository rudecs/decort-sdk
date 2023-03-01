package account

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for delete account
type DeleteRequest struct {
	// ID of account to delete
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`

	// Whether to completely delete the account
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

func (arq DeleteRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	return nil
}

// Delete completes delete an account from the system Returns true if account is deleted or was already deleted or never existed
func (a Account) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/delete"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
