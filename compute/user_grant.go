package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type UserGrantRequest struct {
	ComputeId  uint64 `url:"computeId"`
	Username   string `url:"userName"`
	AccessType string `url:"accesstype"`
}

func (crq UserGrantRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	if crq.Username == "" {
		return errors.New("validation-error: field Username can not be empty")
	}

	if crq.AccessType == "" {
		return errors.New("validation-error: field AccessType can not be empty")
	}
	validator := validators.StringInSlice(crq.AccessType, []string{"R", "RCX", "ARCXDU"})
	if !validator {
		return errors.New("validation-error: field AccessType can be only R, RCX or ARCXDU")
	}

	return nil
}

func (c Compute) UserGrant(ctx context.Context, req UserGrantRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/userGrant"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}
	return result, nil
}
