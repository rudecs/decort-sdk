package grid

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for get backup
type GetBackupRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid"`
}

func (grq GetBackupRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}

	return nil
}

// GetBackup gets platform backup
func (g Grid) GetBackup(ctx context.Context, req GetBackupRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/grid/getBackup"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// GetBackupGET gets platform backup
func (g Grid) GetBackupGET(ctx context.Context, req GetBackupRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/grid/getBackup"

	res, err := g.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
