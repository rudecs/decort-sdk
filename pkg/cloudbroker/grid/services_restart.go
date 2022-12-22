package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restart services
type ServicesRestartRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid"`

	// Node ID
	// Required: true
	NID uint64 `url:"nid"`
}

func (grq ServicesRestartRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.NID == 0 {
		return errors.New("validation-error: field NID must be set")
	}

	return nil
}

// ServicesRestart restarts decort services on the node
func (g Grid) ServicesRestart(ctx context.Context, req ServicesRestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/servicesRestart"

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
