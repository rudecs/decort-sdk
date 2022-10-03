package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type NatRuleDelRequest struct {
	VINSID uint64 `url:"vinsId"`
	RuleID uint64 `url:"ruleId"`
}

func (vrq NatRuleDelRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	if vrq.RuleID == 0 {
		return errors.New("validation-error: field RuleID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) NatRuleDel(ctx context.Context, req NatRuleDelRequest) (bool, error) {
	err := req.Validate()
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
