package vins

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for delete several VINSes
type MassDeleteRequest struct {
	// VINS IDs
	// Required: true
	VINSIDs []uint64 `url:"vinsIds" json:"vinsIds"`

	// Set to true if you want force delete non-empty VINS. Primarily,
	// VINS is considered non-empty if it has VMs connected to it,
	// and force flag will detach them from the VINS being deleted.
	// Otherwise method will return an error
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`

	// Set to true if you want to destroy VINS and all linked resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be restored later
	// within the recycle bins purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq MassDeleteRequest) validate() error {
	if len(vrq.VINSIDs) == 0 {
		return errors.New("validation-error: field VINSIDs must be set")
	}

	return nil
}

// MassDelete start jobs to delete several VINSes
func (v VINS) MassDelete(ctx context.Context, req MassDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/massDelete"

	_, err = v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
