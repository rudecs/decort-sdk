package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RestoreRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq RestoreRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/restore"

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
