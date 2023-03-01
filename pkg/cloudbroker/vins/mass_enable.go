package vins

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for enable several VINSes
type MassEnableRequest struct {
	// VINS IDs
	// Required: true
	VINSIDs []uint64 `url:"vinsIds" json:"vinsIds"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq MassEnableRequest) validate() error {
	if len(vrq.VINSIDs) == 0 {
		return errors.New("validation-error: field VINSIDs must be set")
	}

	return nil
}

// MassEnable start jobs to enable several VINSes
func (v VINS) MassEnable(ctx context.Context, req MassEnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/massEnable"

	_, err = v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
