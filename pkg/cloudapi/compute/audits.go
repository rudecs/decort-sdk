package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AuditsRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq AuditsRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Audits(ctx context.Context, req AuditsRequest) (AuditList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/audits"
	auditListRaw, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	auditList := AuditList{}
	if err := json.Unmarshal(auditListRaw, &auditList); err != nil {
		return nil, err
	}

	return auditList, nil
}
