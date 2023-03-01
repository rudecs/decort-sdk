package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for get info about migration
type MigrateStorageInfoRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq MigrateStorageInfoRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// MigrateStorageInfo gets info about last (include ongoing) storage migration
func (c Compute) MigrateStorageInfo(ctx context.Context, req MigrateStorageInfoRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/migrateStorageInfo"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
