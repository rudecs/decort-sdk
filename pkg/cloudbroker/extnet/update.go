package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update external network
type UpdateRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`

	// New external network name
	// Required: false
	Name string `url:"name,omitempty"`

	// New external network description
	// Required: false
	Description string `url:"desc,omitempty"`
}

func (erq UpdateRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// Update updates external network parameters
func (e ExtNet) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/update"

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
