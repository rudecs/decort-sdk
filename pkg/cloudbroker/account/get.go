package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about account
type GetRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`
}

func (arq GetRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// Get gets information about account
func (a Account) Get(ctx context.Context, req GetRequest) (*RecordAccount, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/account/get"

	info := RecordAccount{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
