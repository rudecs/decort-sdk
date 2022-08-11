package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AuditsRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq AuditsRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) Audits(ctx context.Context, req AuditsRequest, options ...opts.DecortOpts) (AccountAuditsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/audits"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	auditsRaw, err := a.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	audits := AccountAuditsList{}

	err = json.Unmarshal([]byte(auditsRaw), &audits)
	if err != nil {
		return nil, err
	}

	return audits, nil
}
