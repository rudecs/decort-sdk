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
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

func (a Account) ListDisks(ctx context.Context, req ListDisksRequest) (ListDisks, error) {
	err := req.Validate()
	if err != nil {
		return ListDisks{}, err
	}

	url := "/cloudbroker/account/listDisks"

	result := ListDisks{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListDisks{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListDisks{}, err
	}

	return result, nil
}
