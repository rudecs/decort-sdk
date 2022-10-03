package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListTemplatesRequest struct {
	AccountID      uint64 `url:"accountId"`
	IncludeDeleted bool   `url:"includedeleted"`
}

func (arq ListTemplatesRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListTemplates(ctx context.Context, req ListTemplatesRequest) (AccountTemplatesList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listTemplates"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	accountTemplatesList := AccountTemplatesList{}

	err = json.Unmarshal(res, &accountTemplatesList)
	if err != nil {
		return nil, err
	}

	return accountTemplatesList, nil

}
