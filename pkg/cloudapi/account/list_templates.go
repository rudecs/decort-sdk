package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list templates
type ListTemplatesRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Include deleted images
	// Required: false
	IncludeDeleted bool `url:"includedeleted"`
}

func (arq ListTemplatesRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// ListTemplates gets list templates which can be managed by this account
func (a Account) ListTemplates(ctx context.Context, req ListTemplatesRequest) (ListTemplates, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listTemplates"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListTemplates{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
