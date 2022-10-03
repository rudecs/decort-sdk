package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CloneRequest struct {
	ComputeID         uint64 `url:"computeId"`
	Name              string `url:"name"`
	SnapshotTimestamp uint64 `url:"snapshotTimestamp"`
	SnapshotName      string `url:"snapshotName"`
}

func (crq CloneRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Name == "" {
		return errors.New("validation-error: field Name can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Clone(ctx context.Context, req CloneRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/compute/clone"

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
