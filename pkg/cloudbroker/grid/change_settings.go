package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for change grid settings
type ChangeSettingsRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"id" json:"id"`

	// Json data of the new settings will override old data
	// Required: true
	Settings string `url:"settings" json:"settings"`
}

func (grq ChangeSettingsRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.Settings == "" {
		return errors.New("validation-error: field Settings must be set")
	}

	return nil
}

// ChangeSettings changes grid settings
func (g Grid) ChangeSettings(ctx context.Context, req ChangeSettingsRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/changeSettings"

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
