package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get audit records
type AuditsRequest struct {
	// ID of the compute
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq AuditsRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Audits gets audit records for the specified compute object
func (c Compute) Audits(ctx context.Context, req AuditsRequest) (ListAudits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/audits"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
