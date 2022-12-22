package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clear affinity label for compute
type AffinityLabelRemoveRequest struct {
	// IDs of the compute instances
	// Required: true
	ComputeIDs []uint64 `url:"computeIds"`
}

func (crq AffinityLabelRemoveRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// AffinityLabelRemove clear affinity label for compute
func (c Compute) AffinityLabelRemove(ctx context.Context, req AffinityLabelRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/affinityLabelRemove"

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
