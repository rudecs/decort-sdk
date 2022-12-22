package extnet

import (
	"context"
	"net/http"
	"strconv"
)

// GetDefault get default external network ID
func (e ExtNet) GetDefault(ctx context.Context) (uint64, error) {
	url := "/cloudbroker/extnet/getDefault"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
