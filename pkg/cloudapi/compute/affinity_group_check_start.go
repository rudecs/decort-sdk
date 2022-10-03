package compute

import (
	"context"
	"errors"
	"net/http"
)

type AffinityGroupCheckStartRequest struct {
	RGID          uint64 `url:"rgId"`
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityGroupCheckStartRequest) Validate() error {
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel can not be empty")
	}

	return nil
}

func (c Compute) AffinityGroupCheckStart(ctx context.Context, req AffinityGroupCheckStartRequest) (string, error) {
	err := req.Validate()
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
