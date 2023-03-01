package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set affinity label for compute
type AffinityLabelSetRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Affinity group label
	// Required: true
	AffinityLabel string `url:"affinityLabel" json:"affinityLabel"`
}

func (crq AffinityLabelSetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel can not be empty")
	}

	return nil
}

// AffinityLabelSet set affinity label for compute
func (c Compute) AffinityLabelSet(ctx context.Context, req AffinityLabelSetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/affinityLabelSet"

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
