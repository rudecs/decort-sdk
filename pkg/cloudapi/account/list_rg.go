package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListRGRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq ListRGRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListRG(ctx context.Context, req ListRGRequest) (AccountRGList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listRG"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountRGList := AccountRGList{}

	err = json.Unmarshal(res, &accountRGList)
	if err != nil {
		return nil, err
	}

	return accountRGList, nil

}
