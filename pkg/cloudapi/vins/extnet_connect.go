package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for connect external network
type ExtNetConnectRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// External network ID
	// Required: false
	NetID uint64 `url:"netId,omitempty"`

	// Directly set IP address
	// Required: false
	IP string `url:"ip,omitempty"`
}

func (vrq ExtNetConnectRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

// ExtNetConnect connect VINS to external network
func (v VINS) ExtNetConnect(ctx context.Context, req ExtNetConnectRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/extNetConnect"

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
