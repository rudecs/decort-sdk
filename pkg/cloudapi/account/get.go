package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq GetRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) Get(ctx context.Context, req GetRequest) (*AccountWithResources, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/get"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	account := &AccountWithResources{}

	err = json.Unmarshal(res, &account)
	if err != nil {
		return nil, err
	}

	return account, nil

}
