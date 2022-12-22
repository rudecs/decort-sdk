package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clear anti affinity rules
type AntiAffinityRulesClearRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq AntiAffinityRulesClearRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// AntiAffinityRulesClear clear anti affinity rules
func (c Compute) AntiAffinityRulesClear(ctx context.Context, req AntiAffinityRulesClearRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/antiAffinityRulesClear"

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
