package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list snapshots
type SnapshotListRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq SnapshotListRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// SnapshotList gets list compute snapshots
func (c Compute) SnapshotList(ctx context.Context, req SnapshotListRequest) (ListSnapShots, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/snapshotList"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSnapShots{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
