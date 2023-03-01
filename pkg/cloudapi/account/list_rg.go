package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list resource groups
type ListRGRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq ListRGRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// ListRG gets list all resource groups under specified account, accessible by the user
func (a Account) ListRG(ctx context.Context, req ListRGRequest) (ListRG, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listRG"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListRG{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
