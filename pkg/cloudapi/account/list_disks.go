package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list deleted disks
type ListDisksRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq ListDisksRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// ListDisks gets list all currently unattached disks under specified account
func (a Account) ListDisks(ctx context.Context, req ListDisksRequest) (ListDisks, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/listDisks"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDisks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
