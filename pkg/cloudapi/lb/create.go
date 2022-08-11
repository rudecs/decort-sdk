package lb

import (
	"context"
	"errors"
	"strings"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateRequest struct {
	RGID        uint64 `url:"rgId"`
	Name        string `url:"name"`
	ExtnetId    uint64 `url:"extnetId"`
	VinsId      uint64 `url:"vinsId"`
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

	if lbrq.ExtnetId == 0 {
		return errors.New("validation-error: field ExtnetId can not be empty or equal to 0")
	}

	if lbrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	return nil
}

func (l LB) Create(ctx context.Context, req CreateRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/lb/create"
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
