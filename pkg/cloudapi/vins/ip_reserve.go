package vins

import (
	"context"
	"errors"
	"net/http"
)

type IPReserveRequest struct {
	VINSID    uint64 `url:"vinsId"`
	Type      string `url:"type"`
	IPAddr    string `url:"ipAddr,omitempty"`
	MAC       string `url:"mac,omitempty"`
	ComputeID uint64 `url:"computeId,omitempty"`
}

func (vrq IPReserveRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	if vrq.Type == "" {
		return errors.New("validation-error: field Type can not be empty")
	}

	return nil
}

func (v VINS) IPReserve(ctx context.Context, req IPReserveRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/vins/ipReserve"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}
