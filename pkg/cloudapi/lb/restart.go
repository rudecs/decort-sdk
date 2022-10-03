package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RestartRequest struct {
	LBID uint64 `url:"lbId"`
}

func (lbrq RestartRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

func (l LB) Restart(ctx context.Context, req RestartRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/restart"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
