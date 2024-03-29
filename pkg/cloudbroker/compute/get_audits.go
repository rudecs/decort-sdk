package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get compute audits
type GetAuditsRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Reason to action
	// Required: true
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq GetAuditsRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// GetAudits gets compute audits
func (c Compute) GetAudits(ctx context.Context, req GetAuditsRequest) (ListAudits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/getAudits"

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
