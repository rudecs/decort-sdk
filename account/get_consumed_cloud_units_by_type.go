package account

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetConsumedCloudUnitsByTypeRequest struct {
	AccountId uint64 `url:"accountId"`
	CUType    string `url:"cutype"`
}

func (arq GetConsumedCloudUnitsByTypeRequest) Validate() error {
	if arq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	if arq.CUType == "" {
		return errors.New("validation-error: field CUType can not be empty")
	}

	isValid := validators.StringInSlice(arq.CUType, []string{"CU_M", "CU_C", "CU_D", "CU_S", "CU_A", "CU_NO", "CU_I", "CU_NP"})
	if !isValid {
		return errors.New("validation-error: field AccessType can be only CU_M, CU_C, CU_D, CU_S, CU_A, CU_NO, CU_I or CU_NP")
	}

	return nil
}

func (a Account) GetConsumedCloudUnitsByType(ctx context.Context, req GetConsumedCloudUnitsByTypeRequest, options ...opts.DecortOpts) (float64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/account/getConsumedCloudUnitsByType"
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
		return 0, err
	}

	result, err := strconv.ParseFloat(string(res), 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}
