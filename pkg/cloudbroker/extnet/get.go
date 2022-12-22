package extnet

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about external network
type GetRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`
}

func (erq GetRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// Get gets external network details
func (e ExtNet) Get(ctx context.Context, req GetRequest) (*RecordExtNet, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/extnet/get"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordExtNet{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
