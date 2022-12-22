package compute

import (
	"context"
	"net/http"
	"strconv"
)

// RaiseDown starting all computes in "DOWN" tech status
func (c Compute) RaiseDown(ctx context.Context) (bool, error) {
	url := "/cloudbroker/compute/raiseDown"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
