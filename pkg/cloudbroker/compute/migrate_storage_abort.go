package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for abort migration
type MigrateStorageAbortRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq MigrateStorageAbortRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// MigrateStorageAbort abort complex compute migration job
func (c Compute) MigrateStorageAbort(ctx context.Context, req MigrateStorageAbortRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/migrateStorageAbort"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
