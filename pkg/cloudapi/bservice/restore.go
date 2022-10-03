package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RestoreRequest struct {
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq RestoreRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/restore"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
