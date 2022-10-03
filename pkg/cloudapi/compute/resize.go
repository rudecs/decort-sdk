package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ResizeRequest struct {
	ComputeID uint64 `url:"computeId"`
	Force     bool   `url:"force,omitempty"`
	CPU       uint64 `url:"cpu,omitempty"`
	RAM       uint64 `url:"ram,omitempty"`
}

func (crq ResizeRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Resize(ctx context.Context, req ResizeRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/resize"

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
