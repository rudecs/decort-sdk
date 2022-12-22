package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set affinity label for compute
type AffinityLabelSetRequest struct {
	// IDs of the compute instances
	ComputeIDs []uint64 `url:"computeIds"`

	// Affinity group label
	// Required: true
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityLabelSetRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel must be set")
	}

	return nil
}

// AffinityLabelSet set affinity label for compute
func (c Compute) AffinityLabelSet(ctx context.Context, req AffinityLabelSetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/affinityLabelSet"

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
