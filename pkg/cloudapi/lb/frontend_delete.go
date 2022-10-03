package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type FrontendDeleteRequest struct {
	LBID         uint64 `url:"lbId"`
	FrontendName string `url:"frontendName"`
}

func (lbrq FrontendDeleteRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName can not be empty")
	}

	return nil
}

func (l LB) FrontendDelete(ctx context.Context, req FrontendDeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/frontendDelete"

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
