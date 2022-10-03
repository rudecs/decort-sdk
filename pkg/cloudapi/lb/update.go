package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	LBID        uint64 `url:"lbId"`
	Description string `url:"desc"`
}

func (lbrq UpdateRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.Description == "" {
		return errors.New("validation-error: field Description can not be empty")
	}

	return nil
}

func (l LB) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/update"

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
