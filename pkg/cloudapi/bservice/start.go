package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type StartRequest struct {
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq StartRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) Start(ctx context.Context, req StartRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/start"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
