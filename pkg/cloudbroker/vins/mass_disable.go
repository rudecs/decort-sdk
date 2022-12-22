package vins

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for disable several VINSes
type MassDisableRequest struct {
	// VINS IDs
	// Required: true
	VINSIDs []uint64 `url:"vinsIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq MassDisableRequest) validate() error {
	if len(vrq.VINSIDs) == 0 {
		return errors.New("validation-error: field VINSIDs must be set")
	}

	return nil
}

// MassDisable start jobs to disable several VINSes
func (v VINS) MassDisable(ctx context.Context, req MassDisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/massDisable"

	_, err = v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
