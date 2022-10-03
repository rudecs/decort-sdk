package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type StopRequest struct {
	ComputeID uint64 `url:"computeId"`
	Force     bool   `url:"force,omitempty"`
}

func (crq StopRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Stop(ctx context.Context, req StopRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/stop"

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
