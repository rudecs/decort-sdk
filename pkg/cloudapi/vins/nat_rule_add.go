package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type NatRuleAddRequest struct {
	VINSID       uint64 `url:"vinsId"`
	IntIP        string `url:"intIp "`
	IntPort      uint   `url:"intPort"`
	ExtPortStart uint   `url:"extPortStart"`
	ExtPortEnd   uint   `url:"extPortEnd,omitempty"`
	Proto        string `url:"proto"`
}

func (vrq NatRuleAddRequest) Validate() error {
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

func (v VINS) NatRuleAdd(ctx context.Context, req NatRuleAddRequest) (bool, error) {
	err := req.Validate()
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
