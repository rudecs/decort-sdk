package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ConfigResetRequest struct {
	LBID uint64 `url:"lbId"`
}

func (lbrq ConfigResetRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

func (l LB) ConfigReset(ctx context.Context, req ConfigResetRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/configReset"

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
