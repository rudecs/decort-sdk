package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	ServiceID   uint64 `url:"serviceId"`
	Permanently bool   `url:"permanently,omitempty"`
}

func (bsrq DeleteRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/delete"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
