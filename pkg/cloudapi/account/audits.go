package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AuditsRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq AuditsRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) Audits(ctx context.Context, req AuditsRequest) (AccountAuditsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/audits"

	auditsRaw, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	audits := AccountAuditsList{}

	err = json.Unmarshal(auditsRaw, &audits)
	if err != nil {
		return nil, err
	}

	return audits, nil
}
