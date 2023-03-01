package account

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of accounts
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page" json:"page"`

	// Page size
	// Required: false
	Size uint64 `url:"size" json:"size"`
}

// List gets list all accounts the user has access to
func (a Account) List(ctx context.Context, req ListRequest) (ListAccounts, error) {
	url := "/cloudapi/account/list"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAccounts{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
