package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type SnapshotCreateRequest struct {
	ComputeID uint64 `url:"computeId"`
	Label     string `url:"label"`
}

func (crq SnapshotCreateRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Label == "" {
		return errors.New("validation-error: field Label can not be empty")
	}

	return nil
}

func (c Compute) SnapshotCreate(ctx context.Context, req SnapshotCreateRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/compute/snapshotCreate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
