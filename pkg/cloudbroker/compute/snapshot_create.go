package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for create snapshot
type SnapshotCreateRequest struct {
	// ID of the compute instance to create snapshot for
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Text label for snapshot.
	// Must be unique among this compute snapshots
	// Required: true
	Label string `url:"label"`
}

func (crq SnapshotCreateRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.Label == "" {
		return errors.New("validation-error: field Label must be set")
	}

	return nil
}

// SnapshotCreate create compute snapshot
func (c Compute) SnapshotCreate(ctx context.Context, req SnapshotCreateRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/snapshotCreate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
