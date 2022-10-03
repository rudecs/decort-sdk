package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type StartRequest struct {
	ComputeID uint64 `url:"computeId"`
	AltBootID uint64 `url:"altBootId,omitempty"`
}

func (crq StartRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Start(ctx context.Context, req StartRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/start"

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
