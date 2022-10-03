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
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

func (a Account) ListRG(ctx context.Context, req ListRGRequest) (ListRG, error) {
	err := req.Validate()
	if err != nil {
		return ListRG{}, err
	}

	url := "/cloudbroker/account/listRG"

	result := ListRG{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListRG{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListRG{}, err
	}

	return result, nil
}
