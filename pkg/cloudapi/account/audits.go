package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for give list account audits
type AuditsRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId"`
}

func (arq AuditsRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// Audits gets audit records for the specified account object
func (a Account) Audits(ctx context.Context, req AuditsRequest) (ListAudits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/audits"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAudits{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
