package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list VINS
type ListVINSRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq ListVINSRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// ListVINS gets list all ViNSes under specified account, accessible by the user
func (a Account) ListVINS(ctx context.Context, req ListVINSRequest) (ListVINS, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/account/listVins"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListVINS{}, err
	}

	list := ListVINS{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
