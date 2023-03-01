package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable external network
type DisableRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`
}

func (erq DisableRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// Disable disables external network
func (e ExtNet) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/disable"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
