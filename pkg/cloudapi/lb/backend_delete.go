package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type BackendDeleteRequest struct {
	LBID        uint64 `url:"lbId"`
	BackendName string `url:"backendName"`
}

func (lbrq BackendDeleteRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName can not be empty")
	}

	return nil
}

func (l LB) BackendDelete(ctx context.Context, req BackendDeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/backendDelete"

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
