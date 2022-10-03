package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type VNFDevRestartRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq VNFDevRestartRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) VNFDevRestart(ctx context.Context, req VNFDevRestartRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/vnfdevRestart"

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
