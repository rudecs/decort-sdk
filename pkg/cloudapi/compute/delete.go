package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	ComputeID   uint64 `url:"computeId"`
	Permanently bool   `url:"permanently,omitempty"`
	DetachDisks bool   `url:"detachDisks,omitempty"`
}

func (crq DeleteRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/delete"

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
