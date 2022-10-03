package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListVINSRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq ListVINSRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

func (a Account) ListVINS(ctx context.Context, req ListVINSRequest) (ListVINS, error) {
	err := req.Validate()
	if err != nil {
		return ListVINS{}, err
	}

	url := "/cloudbroker/account/listVins"

	result := ListVINS{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListVINS{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListVINS{}, err
	}

	return result, nil
}
