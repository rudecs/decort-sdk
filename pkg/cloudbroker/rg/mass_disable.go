package rg

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for disable several resource groups
type MassDisableRequest struct {
	// IDs of the resource groups
	// Required: true
	RGIDs []uint64 `url:"rgIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq MassDisableRequest) validate() error {
	if len(rgrq.RGIDs) == 0 {
		return errors.New("validation-error: field RGIDs must be set")
	}

	return nil
}

// MassDisable start jobs to disable several resource groups
func (r RG) MassDisable(ctx context.Context, req MassDisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/massDisable"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
