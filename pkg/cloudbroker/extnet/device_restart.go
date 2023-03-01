package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restart VNF device
type DeviceRestartRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`
}

func (erq DeviceRestartRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// DeviceRestart restarts external network VNF device
func (e ExtNet) DeviceRestart(ctx context.Context, req DeviceRestartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/deviceRestart"

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
