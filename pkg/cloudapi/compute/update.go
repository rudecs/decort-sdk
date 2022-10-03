package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	ComputeID   uint64 `url:"computeId"`
	Name        string `url:"name,omitempty"`
	Description string `url:"desc,omitempty"`
}

func (crq UpdateRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/update"

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
