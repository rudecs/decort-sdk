package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type SnapshotListRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq SnapshotListRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) SnapshotList(ctx context.Context, req SnapshotListRequest) (SnapshotList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/snapshotList"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	snapshotList := SnapshotList{}

	err = json.Unmarshal(res, &snapshotList)
	if err != nil {
		return nil, err
	}

	return snapshotList, nil
}
