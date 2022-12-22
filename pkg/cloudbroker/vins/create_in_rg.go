package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create VINS in resource group
type CreateInRGRequest struct {
	// VINS name
	// Required: true
	Name string `url:"name"`

	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Private network IP CIDR
	// Required: false
	IPCIDR string `url:"ipcidr,omitempty"`

	// External network ID
	// Required: false
	ExtNetID uint64 `url:"extNetId,omitempty"`

	// External IP, related only for extNetId >= 0
	// Required: false
	ExtIP string `url:"extIp,omitempty"`

	// Description
	// Required: false
	Description string `url:"desc,omitempty"`

	// Number of pre created reservations
	// Required: false
	PreReservationsNum uint `url:"preReservationsNum,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq CreateInRGRequest) validate() error {
	if vrq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if vrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

// CreateInRG creates VINS in resource group level
func (v VINS) CreateInRG(ctx context.Context, req CreateInRGRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/vins/createInRG"

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
