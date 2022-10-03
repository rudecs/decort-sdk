package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type TagAddRequest struct {
	ComputeID uint64 `url:"computeId"`
	Key       string `url:"key"`
	Value     string `url:"value"`
}

func (crq TagAddRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}

	if crq.Value == "" {
		return errors.New("validation-error: field Value can not be empty")
	}

	return nil
}

func (c Compute) TagAdd(ctx context.Context, req TagAddRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/tagAdd"

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
