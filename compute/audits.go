package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AuditsRequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq AuditsRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Audits(ctx context.Context, req AuditsRequest, options ...opts.DecortOpts) (AuditList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/audits"
	auditListRaw, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	auditList := AuditList{}
	if err := json.Unmarshal(auditListRaw, &auditList); err != nil {
		return nil, err
	}

	return auditList, nil
}
