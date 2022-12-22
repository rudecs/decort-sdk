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
	VINSID uint64 `url:"vinsId"`

	// ID of the rule to delete.
	// Pass -1 to clear all rules at once
	// Required: true
	RuleID uint64 `url:"ruleId"`
}

func (vrq NATRuleDelRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}
	if vrq.RuleID == 0 {
		return errors.New("validation-error: field RuleID can not be empty or equal to 0")
	}

	return nil
}

// NATRuleDel delete NAT (port forwarding) rule on VINS
func (v VINS) NATRuleDel(ctx context.Context, req NATRuleDelRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/natRuleDel"

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
