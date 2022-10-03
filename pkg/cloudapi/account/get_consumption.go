package account

import (
	"context"
	"errors"
	"net/http"
)

type GetConsumtionRequest struct {
	AccountID uint64 `url:"accountId"`
	Start     uint64 `url:"start"`
	End       uint64 `url:"end"`
}

func (arq GetConsumtionRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	if arq.Start == 0 {
		return errors.New("validation-error: field Start can not be empty or equal to 0")
	}

	if arq.End == 0 {
		return errors.New("validation-error: field End can not be empty or equal to 0")
	}

	return nil
}

func (a Account) GetConsumtion(ctx context.Context, req GetConsumtionRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/account/getConsumtion"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}

func (a Account) GetConsumtionGet(ctx context.Context, req GetConsumtionRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/account/getConsumtion"
	prefix := "/cloudapi"

	url = prefix + url
	res, err := a.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}
