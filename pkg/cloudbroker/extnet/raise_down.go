package extnet

import (
	"context"
	"net/http"
	"strconv"
)

// RaiseDown starting all extnets vnfdevs in "DOWN" tech status
func (e ExtNet) RaiseDown(ctx context.Context) (bool, error) {
	url := "/cloudbroker/extnet/raiseDown"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
