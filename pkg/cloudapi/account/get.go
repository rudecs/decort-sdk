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
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq GetRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// Get gets account details
func (a Account) Get(ctx context.Context, req GetRequest) (*RecordAccount, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/get"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordAccount{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil

}
