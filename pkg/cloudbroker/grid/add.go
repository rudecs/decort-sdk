package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for location code
type AddRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid"`

	// Name of the location
	// Required: true
	Name string `url:"name"`

	// Location code typicly used in dns names
	// Required: true
	LocationCode string `url:"locationcode"`
}

func (grq AddRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if grq.LocationCode == "" {
		return errors.New("validation-error: field LocationCode must be set")
	}

	return nil
}

// Add location code (e.g. DNS name of this grid)
func (g Grid) Add(ctx context.Context, req AddRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/add"

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
