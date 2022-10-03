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
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListVINS(ctx context.Context, req ListVINSRequest) (AccountVINSList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listVins"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountVINSList := AccountVINSList{}

	err = json.Unmarshal(res, &accountVINSList)
	if err != nil {
		return nil, err
	}

	return accountVINSList, nil

}
