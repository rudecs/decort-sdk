package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start VNF devices
type VNFDevStartRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq VNFDevStartRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: fiels VINSID must be set")
	}

	return nil
}

// VNFDevStart starts VINSes primary VNF device
func (v VINS) VNFDevStart(ctx context.Context, req VNFDevStartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/vnfdevStart"

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
