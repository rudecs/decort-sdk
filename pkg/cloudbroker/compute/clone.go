package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for clone compute instance
type CloneRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Name of the clone
	// Required: true
	Name string `url:"name" json:"name"`

	// Timestamp of the parent's snapshot to create clone from
	// Required: false
	SnapshotTimestamp uint64 `url:"snapshotTimestamp" json:"snapshotTimestamp"`

	// Name of the parent's snapshot to create clone from
	// Required: false
	SnapshotName string `url:"snapshotName" json:"snapshotName"`

	// Reason to clone
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq CloneRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// Clone clones compute instance
func (c Compute) Clone(ctx context.Context, req CloneRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/compute/clone"

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
