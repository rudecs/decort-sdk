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
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Internal IP address to apply this rule to
	// Required: true
	IntIP string `url:"intIp " json:"intIp "`

	// Internal IP port number to use for this rule
	// Required: true
	IntPort uint `url:"intPort" json:"intPort"`

	// External IP start port to use for this rule
	// Required: true
	ExtPortStart uint `url:"extPortStart" json:"extPortStart"`

	// External IP end port to use for this rule
	// Required: false
	ExtPortEnd uint `url:"extPortEnd,omitempty" json:"extPortEnd,omitempty"`

	// IP protocol type
	// Should be one of:
	//	- "tcp"
	//	- "udp"
	// Required: false
	Proto string `url:"proto,omitempty" json:"proto,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq NATRuleAddRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}
	if vrq.IntIP == "" {
		return errors.New("validation-error: field IntIP must be set")
	}
	if vrq.IntPort == 0 {
		return errors.New("validation-error: field IntPort must be set")
	}
	if vrq.ExtPortStart == 0 {
		return errors.New("validation-error: field ExtPortStart must be set")
	}

	return nil
}

// NATRuleAdd create NAT (port forwarding) rule on VINS
func (v VINS) NATRuleAdd(ctx context.Context, req NATRuleAddRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/vins/natRuleAdd"

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
