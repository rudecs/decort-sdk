package lb

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type CreateRequest struct {
	RGID        uint64 `url:"rgId"`
	Name        string `url:"name"`
	ExtNetID    uint64 `url:"extnetId"`
	VINSID      uint64 `url:"vinsId"`
	Start       bool   `url:"start"`
	Description string `url:"desc,omitempty"`
}

func (lbrq CreateRequest) Validate() error {
	if lbrq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}

	if lbrq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if lbrq.ExtNetID == 0 {
		return errors.New("validation-error: field ExtNetID can not be empty or equal to 0")
	}

	if lbrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (l LB) Create(ctx context.Context, req CreateRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/lb/create"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(res), "\"", ""), nil
}
