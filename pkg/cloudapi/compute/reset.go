package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ResetRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq ResetRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Reset(ctx context.Context, req ResetRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/reset"

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
