package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create NAT rules
type NATRuleAddRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Internal IP address to apply this rule to
	// Required: true
	IntIP string `url:"intIp "`

	// Internal IP port number to use for this rule
	// Required: true
	IntPort uint `url:"intPort"`

	// External IP start port to use for this rule
	// Required: true
	ExtPortStart uint `url:"extPortStart"`

	// External IP end port to use for this rule
	// Required: false
	ExtPortEnd uint `url:"extPortEnd,omitempty"`

	// IP protocol type
	// Should be one of:
	//	- "tcp"
	//	- "udp"
	// Required: false
	Proto string `url:"proto,omitempty"`
}

func (vrq NATRuleAddRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}
	if vrq.IntIP == "" {
		return errors.New("validation-error: field IntIP can not be empty")
	}
	if vrq.IntPort == 0 {
		return errors.New("validation-error: field IntPort can not be empty or equal to 0")
	}
	if vrq.ExtPortStart == 0 {
		return errors.New("validation-error: field ExtPortStart can not be empty or equal to 0")
	}

	return nil
}

// NATRuleAdd create NAT (port forwarding) rule on VINS
func (v VINS) NATRuleAdd(ctx context.Context, req NATRuleAddRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/natRuleAdd"

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
