package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for destroy
type DestroyRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`
}

func (erq DestroyRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// Destroy destroys external network
func (e ExtNet) Destroy(ctx context.Context, req DestroyRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/destroy"

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
