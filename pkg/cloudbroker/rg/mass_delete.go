package rg

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for delete several resource groups
type MassDeleteRequest struct {
	// IDs of the resource groups
	// Required: true
	RGIDs []uint64 `url:"rgIds"`

	// Set to true if you want force delete non-empty resource groups
	// Required: false
	Force bool `url:"force,omitempty"`

	// Set to true if you want to destroy resource group and all linked
	// resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be
	// restored later within recycle bins purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq MassDeleteRequest) validate() error {
	if len(rgrq.RGIDs) == 0 {
		return errors.New("validation-error: field RGIDs must be set")
	}

	return nil
}

// MassDelete starts jobs to delete several resource groups
func (r RG) MassDelete(ctx context.Context, req MassDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/massDelete"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
