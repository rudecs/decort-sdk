package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create VINS in account
type CreateInAccountRequest struct {
	// VINS name
	// Required: true
	Name string `url:"name"`

	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Grid ID
	// Required: false
	GID uint64 `url:"gid,omitempty"`

	// Private network IP CIDR
	// Required: false
	IPCIDR string `url:"ipcidr,omitempty"`

	// Description
	// Required: false
	Description string `url:"desc,omitempty"`

	// Number of pre created reservations
	// Required: false
	PreReservationsNum uint64 `url:"preReservationsNum,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq CreateInAccountRequest) validate() error {
	if vrq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if vrq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// CreateInAccount creates VINS in account level
func (v VINS) CreateInAccount(ctx context.Context, req CreateInAccountRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/vins/createInAccount"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
