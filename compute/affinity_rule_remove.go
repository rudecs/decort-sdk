package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AffinityRuleRemoveRequest struct {
	ComputeId uint64 `url:"computeId"`
	Topology  string `url:"topology"`
	Policy    string `url:"policy"`
	Mode      string `url:"mode"`
	Key       string `url:"key"`
	Value     string `url:"value"`
}

func (crq AffinityRuleRemoveRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}
	if crq.Topology == "" {
		return errors.New("validation-error: field Topology can not be empty")
	}

	validator := validators.StringInSlice(crq.Topology, []string{"compute", "node"})
	if !validator {
		return errors.New("validation-error: field Topology can be only compute or node")
	}

	if crq.Policy == "" {
		return errors.New("validation-error: field Policy can not be empty")
	}

	validator = validators.StringInSlice(crq.Policy, []string{"RECOMMENDED", "REQUIRED"})
	if !validator {
		return errors.New("validation-error: field Policy can be only RECOMMENDED or REQUIRED")
	}

	if crq.Mode == "" {
		return errors.New("validation-error: field Mode can not be empty")
	}

	validator = validators.StringInSlice(crq.Mode, []string{"EQ", "NE", "ANY"})
	if !validator {
		return errors.New("validation-error: field Mode can be only EQ, NE or ANY")
	}

	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}

	if crq.Value == "" {
		return errors.New("validation-error: field Value can not be empty")
	}

	return nil
}

func (c Compute) AffinityRuleRemove(ctx context.Context, req AffinityRuleRemoveRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/affinityRuleRemove"
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
