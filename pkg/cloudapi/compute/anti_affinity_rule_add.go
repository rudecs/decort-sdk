package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type AntiAffinityRuleAddRequest struct {
	ComputeID uint64 `url:"computeId"`
	Topology  string `url:"topology"`
	Policy    string `url:"policy"`
	Mode      string `url:"mode"`
	Key       string `url:"key"`
	Value     string `url:"value"`
}

func (crq AntiAffinityRuleAddRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
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

func (c Compute) AntiAffinityRuleAdd(ctx context.Context, req AntiAffinityRuleAddRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/antiAffinityRuleAdd"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
