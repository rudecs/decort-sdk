package vins

import (
	"context"
	"net/http"
	"strconv"
)

// RaiseDown starting all VINSes VNFDevs in "DOWN" tech status
func (v VINS) RaiseDown(ctx context.Context) (bool, error) {
	url := "/cloudbroker/vins/raiseDown"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
