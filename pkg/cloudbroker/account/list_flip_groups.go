package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list FLIPGroups
type ListFLIPGroupsRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq ListFLIPGroupsRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// ListFLIPGroups gets list all FLIPGroups under specified account, accessible by the user
func (a Account) ListFLIPGroups(ctx context.Context, req ListFLIPGroupsRequest) (ListFLIPGroups, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/account/listFlipGroups"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListFLIPGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
