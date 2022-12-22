package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for cleanup resources after finished migration
type MigrateStorageCleanUpRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq MigrateStorageCleanUpRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// MigrateStorageCleanUp cleanup resources after finished (success of failed) complex compute migration.
// If the migration was successful, then old disks will be removed, else new (target) disks will be removed.
// Do it wisely!
func (c Compute) MigrateStorageCleanUp(ctx context.Context, req MigrateStorageCleanUpRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/migrateStorageCleanup"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
