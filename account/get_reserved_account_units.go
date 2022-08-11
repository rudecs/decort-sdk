package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetReservedAccountUnitsRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (arq GetReservedAccountUnitsRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (a Account) GetReservedAccountUnits(ctx context.Context, req GetReservedAccountUnitsRequest, options ...opts.DecortOpts) (*ResourceLimits, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/account/getReservedAccountUnits"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := a.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	rl := &ResourceLimits{}

	err = json.Unmarshal(res, &rl)
	if err != nil {
		return nil, err
	}

	return rl, nil

}
