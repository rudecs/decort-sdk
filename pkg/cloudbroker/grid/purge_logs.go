package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for purge logs
type PurgeLogsRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid"`

	// Age of the records to remove, e.g. -1h for records older than 1 hour, -1w - one week, etc
	// Required: true
	Age string `url:"age"`
}

func (grq PurgeLogsRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.Age == "" {
		return errors.New("validation-error: field Age must be set")
	}

	return nil
}

// PurgeLogs clear Log and ECO records that are older than the specified age.
// By default, records older than one week are removed
func (g Grid) PurgeLogs(ctx context.Context, req PurgeLogsRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/purgeLogs"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
