package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable VINS
type DisableRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq DisableRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// Disable disables VINS by ID
func (v VINS) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/disable"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
