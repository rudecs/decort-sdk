package extnet

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for grant access
type AccessAddRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`

	// Account ID
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (erq AccessAddRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}
	if erq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// AccessAdd grant access to external network for account ID
func (e ExtNet) AccessAdd(ctx context.Context, req AccessAddRequest) ([]uint64, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/extnet/accessAdd"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]uint64, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
