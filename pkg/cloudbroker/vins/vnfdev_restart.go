package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for reboot VINSes primary VNF device
type VNFDevRestartRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Reason for action
	// Required: fal
	Reason string `url:"reason,omitempty"`
}

func (vrq VNFDevRestartRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: fiels VINSID must be set")
	}

	return nil
}

// VNFDevRestart reboot VINSes primary VNF device
func (v VINS) VNFDevRestart(ctx context.Context, req VNFDevRestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/vnfdevRestart"

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
