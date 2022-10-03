package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type GetConsumedCloudUnitsByTypeRequest struct {
	AccountID uint64 `url:"accountId"`
	CUType    string `url:"cutype"`
}

func (arq GetConsumedCloudUnitsByTypeRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
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

func (a Account) GetConsumedCloudUnitsByType(ctx context.Context, req GetConsumedCloudUnitsByTypeRequest) (float64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/account/getConsumedCloudUnitsByType"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseFloat(string(res), 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}
