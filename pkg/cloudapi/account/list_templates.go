package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListTemplatesRequest struct {
	AccountId      uint64 `url:"accountId"`
	IncludeDeleted bool   `url:"includedeleted"`
}

func (arq ListTemplatesRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) ListTemplates(ctx context.Context, req ListTemplatesRequest, options ...opts.DecortOpts) (AccountTemplatesList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/listTemplates"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := a.client.DecortApiCall(ctx, typed.POST, url, req)
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
