package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RestoreRequest struct {
	LBID uint64 `url:"lbId"`
}

func (lbrq RestoreRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

func (l LB) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/restore"

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
