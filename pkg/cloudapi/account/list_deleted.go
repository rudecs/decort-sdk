package account

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (a Account) ListDeleted(ctx context.Context, req ListDeletedRequest) (AccountCloudApiList, error) {
	url := "/cloudapi/account/listDeleted"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountList := AccountCloudApiList{}

	err = json.Unmarshal(res, &accountList)
	if err != nil {
		return nil, err
	}

	return accountList, nil

}
