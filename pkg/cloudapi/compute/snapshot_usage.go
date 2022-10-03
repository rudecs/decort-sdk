package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type SnapshotUsageRequest struct {
	ComputeID uint64 `url:"computeId"`
	Label     string `url:"label,omitempty"`
}

func (crq SnapshotUsageRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) SnapshotUsage(ctx context.Context, req SnapshotUsageRequest) (SnapshotUsageList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/snapshotUsage"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	snapshotUsage := SnapshotUsageList{}

	err = json.Unmarshal(res, &snapshotUsage)
	if err != nil {
		return nil, err
	}

	return snapshotUsage, nil
}
