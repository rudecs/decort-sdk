package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disconnect VINS from external network
type ExtNetDisconnectRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq ExtNetDisconnectRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// ExtNetDisconnect disconnect VINS from external network
func (v VINS) ExtNetDisconnect(ctx context.Context, req ExtNetDisconnectRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/extNetDisconnect"

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
