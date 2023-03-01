package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get dict of computes
type AffinityRelationsRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq AffinityRelationsRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// AffinityRelations gets dict of computes divided by affinity and anti affinity rules
func (c Compute) AffinityRelations(ctx context.Context, req AffinityRelationsRequest) (*RecordAffinityRelations, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/affinityRelations"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordAffinityRelations{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
