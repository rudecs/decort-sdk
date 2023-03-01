package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clear affinity label for compute
type AffinityLabelRemoveRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq AffinityLabelRemoveRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// AffinityLabelRemove clear affinity label for compute
func (c Compute) AffinityLabelRemove(ctx context.Context, req AffinityLabelRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/affinityLabelRemove"

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
