package vins

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type NatRuleAddRequest struct {
	VinsId       uint64 `url:"vinsId"`
	IntIP        string `url:"intIp "`
	IntPort      uint   `url:"intPort"`
	ExtPortStart uint   `url:"extPortStart"`
	ExtPortEnd   uint   `url:"extPortEnd,omitempty"`
	Proto        string `url:"proto"`
}

func (vrq NatRuleAddRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
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

func (v Vins) NatRuleAdd(ctx context.Context, req NatRuleAddRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/vins/natRuleAdd"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := v.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, nil
	}

	return result, nil

}
