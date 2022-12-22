package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get compute snapshot real size on storage
type SnapshotUsageRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Specify to show usage exact for this snapshot.
	// Leave empty for get usage for all compute snapshots
	// Required: false
	Label string `url:"label,omitempty"`
}

func (crq SnapshotUsageRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// SnapshotUsage Get compute snapshot real size on storage.
// Always returns list of json objects, and first json object contains summary about all related
// snapshots.
func (c Compute) SnapshotUsage(ctx context.Context, req SnapshotUsageRequest) (ListSnapshotUsage, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/snapshotUsage"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSnapshotUsage{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
