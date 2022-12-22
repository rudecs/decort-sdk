package grid

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get grid details
type GetRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gridId"`
}

func (grq GetRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}

	return nil
}

// Get gets information about grid by ID
func (g Grid) Get(ctx context.Context, req GetRequest) (*RecordGrid, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/grid/get"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordGrid{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
