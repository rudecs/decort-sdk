package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

func (r RG) Audits(ctx context.Context, req AuditsRequest) (AuditList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/audits"
	auditListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	auditList := AuditList{}
	if err := json.Unmarshal(auditListRaw, &auditList); err != nil {
		return nil, err
	}

	return auditList, nil
}
