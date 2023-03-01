package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for check virtual machine
type CheckVMsRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid"`
}

func (grq CheckVMsRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}

	return nil
}

// CheckVMs run checkvms jumpscript
func (g Grid) CheckVMs(ctx context.Context, req CheckVMsRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/checkVMs"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
