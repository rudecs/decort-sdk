package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for a get list compute instances
type ListComputesRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`
}

func (arq ListComputesRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// ListComputes gets list all compute instances under specified account, accessible by the user
func (a Account) ListComputes(ctx context.Context, req ListComputesRequest) (ListComputes, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listComputes"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
