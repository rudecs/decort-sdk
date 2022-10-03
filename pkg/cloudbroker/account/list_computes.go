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
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

func (a Account) ListComputes(ctx context.Context, req ListComputesRequest) (ListComputes, error) {
	err := req.Validate()
	if err != nil {
		return ListComputes{}, err
	}

	url := "/cloudbroker/account/listComputes"

	result := ListComputes{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListComputes{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListComputes{}, err
	}

	return result, nil
}
