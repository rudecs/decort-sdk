package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete NAT rule
type NATRuleDelRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// ID of the rule to delete.
	// Pass -1 to clear all rules at once
	// Required: true
	RuleID uint64 `url:"ruleId" json:"ruleId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq NATRuleDelRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}
	if vrq.RuleID == 0 {
		return errors.New("validation-error: field RuleID must be set")
	}

	return nil
}

// NATRuleDel delete NAT (port forwarding) rule on VINS
func (v VINS) NATRuleDel(ctx context.Context, req NATRuleDelRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/natRuleDel"

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
