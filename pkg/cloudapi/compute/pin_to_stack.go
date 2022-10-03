package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type PinToStackRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq PinToStackRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) PinToStack(ctx context.Context, req PinToStackRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/compute/pinToStack"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
