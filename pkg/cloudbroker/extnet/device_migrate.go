package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for migrate VNF
type DeviceMigrateRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`
}

func (erq DeviceMigrateRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// DeviceMigrate migrate external network VNF device
func (e ExtNet) DeviceMigrate(ctx context.Context, req DeviceMigrateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/deviceMigrate"

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
