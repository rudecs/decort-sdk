package lb

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type FrontendBindDeleteRequest struct {
	LBID         uint64 `url:"lbId"`
	FrontendName string `url:"frontendName"`
	BindingName  string `url:"bindingName"`
}

func (lbrq FrontendBindDeleteRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	if lbrq.FrontendName == "" {
		return errors.New("validation-error: field FrontendName can not be empty")
	}

	if lbrq.BindingName == "" {
		return errors.New("validation-error: field BindingName can not be empty")
	}

	return nil
}

func (l LB) FrontendBindDelete(ctx context.Context, req FrontendBindDeleteRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/lb/frontendBindDelete"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(res), "\"", ""), nil
}
