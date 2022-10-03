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
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

func (a Account) ListFlipGroups(ctx context.Context, req ListFlipGroupsRequest) (ListFlipGroups, error) {
	err := req.Validate()
	if err != nil {
		return ListFlipGroups{}, err
	}

	url := "/cloudbroker/account/listFlipGroups"

	result := ListFlipGroups{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListFlipGroups{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListFlipGroups{}, err
	}

	return result, nil
}
