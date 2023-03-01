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
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// External network ID
	// Required: true
	NetID uint64 `url:"netId" json:"netId"`

	// Directly set IP address
	// Required: false
	IP string `url:"ip,omitempty" json:"ip,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq ExtNetConnectRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}
	if vrq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// ExtNetConnect connect VINS to external network
func (v VINS) ExtNetConnect(ctx context.Context, req ExtNetConnectRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/extNetConnect"

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
