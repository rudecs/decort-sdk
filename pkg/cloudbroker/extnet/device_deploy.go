package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for deploy network device
type DeviceDeployRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`
}

func (erq DeviceDeployRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// DeviceDeploy deploys network device for external network (make not virtual, "physical")
func (e ExtNet) DeviceDeploy(ctx context.Context, req DeviceDeployRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/deviceDeploy"

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
