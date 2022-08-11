package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AuditsRequest struct {
	RGID uint64 `url:"rgId"`
}

func (rgrq AuditsRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) Audits(ctx context.Context, req AuditsRequest, options ...opts.DecortOpts) (AuditList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/audits"
	auditListRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	auditList := AuditList{}
	if err := json.Unmarshal(auditListRaw, &auditList); err != nil {
		return nil, err
	}

	return auditList, nil
}
