package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DiskDetachRequest struct {
	ComputeID uint64 `url:"computeId"`
	DiskID    uint64 `url:"diskId"`
}

func (crq DiskDetachRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) DiskDetach(ctx context.Context, req DiskDetachRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/diskDetach"

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
