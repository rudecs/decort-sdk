package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for remove network device
type DeviceRemoveRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`
}

func (erq DeviceRemoveRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// DeviceRemove removes network device of external network (make it virtual, not "physical")
func (e ExtNet) DeviceRemove(ctx context.Context, req DeviceRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/deviceRemove"

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
