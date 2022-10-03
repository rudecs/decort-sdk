package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type TagRemoveRequest struct {
	ComputeID uint64 `url:"computeId"`
	Key       string `url:"key"`
}

func (crq TagRemoveRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}

	return nil
}

func (c Compute) TagRemove(ctx context.Context, req TagRemoveRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/tagRemove"

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
