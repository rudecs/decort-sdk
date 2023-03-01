package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list of NAT rules
type NATRuleListRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq NATRuleListRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// NATRuleList gets list of NAT (port forwarding) rules
func (v VINS) NATRuleList(ctx context.Context, req NATRuleListRequest) (ListNATRules, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/vins/natRuleList"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListNATRules{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
