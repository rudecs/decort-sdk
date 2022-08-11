package vins

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AuditsRequest struct {
	VinsId uint64 `url:"vinsId"`
}

func (vrq AuditsRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) Audits(ctx context.Context, req AuditsRequest, options ...opts.DecortOpts) (VinsAuditsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/vins/audits"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	auditsRaw, err := v.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	audits := VinsAuditsList{}

	err = json.Unmarshal([]byte(auditsRaw), &audits)
	if err != nil {
		return nil, err
	}

	return audits, nil
}
