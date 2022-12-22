package rg

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for enable several resource groups
type MassEnableRequest struct {
	// IDs of the resource groups
	// Required: true
	RGIDs []uint64 `url:"rgIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq MassEnableRequest) validate() error {
	if len(rgrq.RGIDs) == 0 {
		return errors.New("validation-error: field RGIDs must be set")
	}

	return nil
}

// MassEnable start jobs to enable several resource groups
func (r RG) MassEnable(ctx context.Context, req MassEnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/massEnable"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
