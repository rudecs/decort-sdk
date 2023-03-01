package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clear anti affinity rules
type AntiAffinityRulesClearRequest struct {
	// IDs of the compute instances
	// Required: true
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds"`
}

func (crq AntiAffinityRulesClearRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}

	return nil
}

// AntiAffinityRulesClear clear anti affinity rules
func (c Compute) AntiAffinityRulesClear(ctx context.Context, req AntiAffinityRulesClearRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/antiAffinityRulesClear"

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
