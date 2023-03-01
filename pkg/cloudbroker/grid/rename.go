package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for rename grid
type RenameRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid"`

	// New name
	// Required: true
	Name string `url:"Name" json:"Name"`
}

func (grq RenameRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// Rename renames a grid
func (g Grid) Rename(ctx context.Context, req RenameRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/rename"

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
