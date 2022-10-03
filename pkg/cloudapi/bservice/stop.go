package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type StopRequest struct {
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq StopRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) Stop(ctx context.Context, req StopRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/stop"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
