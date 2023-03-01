package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set external network as default
type SetDefaultRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`
}

func (erq SetDefaultRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// SetDefault sets external network as default for platform
func (e ExtNet) SetDefault(ctx context.Context, req SetDefaultRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/setDefault"

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
