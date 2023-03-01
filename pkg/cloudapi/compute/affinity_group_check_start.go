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
	RGID uint64 `url:"rgId" json:"rgId"`

	// Affinity group label
	// Required: true
	AffinityLabel string `url:"affinityLabel" json:"affinityLabel"`
}

func (crq AffinityGroupCheckStartRequest) validate() error {
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel can not be empty")
	}

	return nil
}

// AffinityGroupCheckStart check all computes with current affinity label can start
func (c Compute) AffinityGroupCheckStart(ctx context.Context, req AffinityGroupCheckStartRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/compute/affinityGroupCheckStart"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
