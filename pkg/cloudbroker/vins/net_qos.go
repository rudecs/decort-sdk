package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update all VINS interfaces QOS
type NetQOSRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Internal traffic, kbit
	// Required: false
	IngressRate uint64 `url:"ingress_rate,omitempty" json:"ingress_rate,omitempty"`

	// Internal traffic burst, kbit
	// Required: false
	IngressBirst uint64 `url:"ingress_birst,omitempty" json:"ingress_birst,omitempty"`

	// External traffic rate, kbit
	// Required: false
	EgressRate uint64 `url:"egress_rate,omitempty" json:"egress_rate,omitempty"`
}

func (vrq NetQOSRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// NetQOS update all VINS interfaces QOS
func (v VINS) NetQOS(ctx context.Context, req NetQOSRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/netQos"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
