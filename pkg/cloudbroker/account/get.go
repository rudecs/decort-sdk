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
		return errors.New("validation-error: field AccountID must be set")
	}
	return nil
}

func (a Account) Get(ctx context.Context, req GetRequest) (GetResponse, error) {
	err := req.Validate()
	if err != nil {
		return GetResponse{}, err
	}

	url := "/cloudbroker/account/get"

	result := GetResponse{}
	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return GetResponse{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return GetResponse{}, err
	}

	return result, nil
}
