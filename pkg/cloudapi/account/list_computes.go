package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListComputesRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq ListComputesRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListComputes(ctx context.Context, req ListComputesRequest) (AccountComputesList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listComputes"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountComputesList := AccountComputesList{}

	err = json.Unmarshal(res, &accountComputesList)
	if err != nil {
		return nil, err
	}

	return accountComputesList, nil

}
