package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListDisksRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq ListDisksRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListDisks(ctx context.Context, req ListDisksRequest) (AccountDisksList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listDisks"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountDisksList := AccountDisksList{}

	err = json.Unmarshal(res, &accountDisksList)
	if err != nil {
		return nil, err
	}

	return accountDisksList, nil

}
