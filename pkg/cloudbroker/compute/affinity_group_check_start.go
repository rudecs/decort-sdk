package compute

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for check all computes with current affinity label can start
type AffinityGroupCheckStartRequest struct {
	// ID of the resource group
	// Required: true
	RGID uint64 `url:"rgId"`

	// Affinity group label
	// Required: true
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityGroupCheckStartRequest) validate() error {
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel must be set")
	}

	return nil
}

// AffinityGroupCheckStart check all computes with current affinity label can start
func (c Compute) AffinityGroupCheckStart(ctx context.Context, req AffinityGroupCheckStartRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/affinityGroupCheckStart"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
