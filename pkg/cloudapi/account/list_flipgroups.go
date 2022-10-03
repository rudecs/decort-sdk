package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListFlipGroupsRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq ListFlipGroupsRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListFlipGroups(ctx context.Context, req ListFlipGroupsRequest) (AccountFlipGroupsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listFlipGroups"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountFlipGroupsList := AccountFlipGroupsList{}

	err = json.Unmarshal(res, &accountFlipGroupsList)
	if err != nil {
		return nil, err
	}

	return accountFlipGroupsList, nil

}
