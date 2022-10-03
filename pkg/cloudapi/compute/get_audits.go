package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetAuditsRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq GetAuditsRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) GetAudits(ctx context.Context, req GetAuditsRequest) (AuditShortList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/getAudits"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	auditsList := AuditShortList{}
	err = json.Unmarshal(res, &auditsList)
	if err != nil {
		return nil, err
	}

	return auditsList, nil
}
