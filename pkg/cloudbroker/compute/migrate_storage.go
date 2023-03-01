package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for migration
type MigrateStorageRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// SEP ID to migrate disks
	// Required: true
	SEPID uint64 `url:"sepId" json:"sepId"`

	// SEP pool name to migrate disks
	// Required: true
	PoolName string `url:"poolName" json:"poolName"`

	// Target stack ID
	// Required: true
	StackID uint64 `url:"stackId" json:"stackId"`

	// Async API call
	// Required: true
	Sync bool `url:"sync" json:"sync"`
}

func (crq MigrateStorageRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if crq.PoolName == "" {
		return errors.New("validation-error: field PoolName must be set")
	}
	if crq.StackID == 0 {
		return errors.New("validation-error: field StackID must be set")
	}

	return nil
}

// MigrateStorage gets complex compute migration
// Compute will be migrated to specified stack, and compute disks will
// be migrated to specified SEP to specified pool.
// This action can take up to 84 hours
func (c Compute) MigrateStorage(ctx context.Context, req MigrateStorageRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/migrateStorage"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
