package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type AffinityLabelSetRequest struct {
	ComputeID     uint64 `url:"computeId"`
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityLabelSetRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel can not be empty")
	}

	return nil
}

func (c Compute) AffinityLabelSet(ctx context.Context, req AffinityLabelSetRequest) (bool, error) {
	err := req.Validate()
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
