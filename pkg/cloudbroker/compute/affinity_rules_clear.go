package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clear affinity rules
type AffinityRulesClearRequest struct {
	// IDs of the compute instances
	// Required: true
	ComputeIDs []uint64 `url:"computeIds"`
}

func (crq AffinityRulesClearRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// AffinityRulesClear clear affinity rules
func (c Compute) AffinityRulesClear(ctx context.Context, req AffinityRulesClearRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/affinityRulesClear"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
