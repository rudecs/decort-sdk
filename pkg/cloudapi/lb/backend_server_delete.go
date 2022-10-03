package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type BackendServerDeleteRequest struct {
	LBID        uint64 `url:"lbId"`
	BackendName string `url:"backendName"`
	ServerName  string `url:"serverName"`
}

func (lbrq BackendServerDeleteRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName can not be empty")
	}

	if lbrq.ServerName == "" {
		return errors.New("validation-error: field ServerName can not be empty")
	}

	return nil
}

func (l LB) BackendServerDelete(ctx context.Context, req BackendServerDeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/backendServerDelete"

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
