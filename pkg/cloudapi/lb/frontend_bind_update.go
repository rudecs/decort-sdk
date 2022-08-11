package lb

import (
	"context"
	"errors"
	"strings"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type FrontendBindUpdateRequest struct {
	LBID           uint64 `url:"lbId"`
	FrontendName   string `url:"frontendName"`
	BindingName    string `url:"bindingName"`
	BindingAddress string `url:"bindingAddress,omitempty"`
	BindingPort    uint   `url:"bindingPort,omitempty"`
}

func (lbrq FrontendBindUpdateRequest) Validate() error {
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

func (l LB) FrontendBindUpdate(ctx context.Context, req FrontendBindUpdateRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/lb/frontendBindingUpdate"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := l.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(res), "\"", ""), nil
}
