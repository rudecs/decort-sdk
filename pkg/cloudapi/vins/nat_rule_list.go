package vins

import (
	"context"
	"errors"
	"net/http"
)

type NatRuleListRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq NatRuleListRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) NatRuleList(ctx context.Context, req NatRuleListRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/vins/natRuleList"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}
