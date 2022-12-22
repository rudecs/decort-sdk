package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for stop VNF devices
type VNFDevStopRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Reason for action
	// Required: true
	Reason string `url:"reason,omitempty"`
}

func (vrq VNFDevStopRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: fiels VINSID must be set")
	}

	return nil
}

// VNFDevStop stop VINSes primary VNF device
func (v VINS) VNFDevStop(ctx context.Context, req VNFDevStopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/vnfdevStop"

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
