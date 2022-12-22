package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for enable external network
type EnableRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`
}

func (erq EnableRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// Enable enables external networks
func (e ExtNet) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/enable"

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
