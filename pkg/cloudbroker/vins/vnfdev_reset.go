package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for reset VNF device
type VNFDevResetRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq VNFDevResetRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: fiels VINSID must be set")
	}

	return nil
}

// VNFDevReset reset VINSes primary VNF device
func (v VINS) VNFDevReset(ctx context.Context, req VNFDevResetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/vnfdevReset"

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
