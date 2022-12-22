package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for execute script
type ExecuteMaintenanceScriptRequest struct {
	// Grid (platform) ID
	// Required: true
	GID string `url:"gid"`

	// Type of nodes you want to apply the action on
	// Required: true
	NodesType string `url:"nodestype"`

	// The script you want to run
	// Required: true
	Script string `url:"script"`
}

func (grq ExecuteMaintenanceScriptRequest) validate() error {
	if grq.GID == "" {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.NodesType == "" {
		return errors.New("validation-error: field NodesType must be set")
	}
	if grq.Script == "" {
		return errors.New("validation-error: field Script must be set")
	}

	return nil
}

// ExecuteMaintenanceScript executes maintenance script
func (g Grid) ExecuteMaintenanceScript(ctx context.Context, req ExecuteMaintenanceScriptRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/executeMaintenanceScript"

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
