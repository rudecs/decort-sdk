package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetAuditsRequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq GetAuditsRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) GetAudits(ctx context.Context, req GetAuditsRequest, options ...opts.DecortOpts) (AuditShortList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/getAudits"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
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
