package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for calculate the currently consumed cloud units of the specified type for all cloudspaces and resource groups in the account
type GetConsumedCloudUnitsByTypeRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Cloud unit resource type
	// Required: true
	CUType string `url:"cutype"`
}

func (arq GetConsumedCloudUnitsByTypeRequest) validate() error {
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

// GetConsumedCloudUnitsByType calculates the currently consumed cloud units of the specified type for all cloudspaces
// and resource groups in the account.
// Possible types of cloud units are include:
//
//   - CU_M: returns consumed memory in MB
//   - CU_C: returns number of virtual cpu cores
//   - CU_D: returns consumed virtual disk storage in GB
//   - CU_S: returns consumed primary storage (NAS) in TB
//   - CU_A: returns consumed secondary storage (Archive) in TB
//   - CU_NO: returns sent/received network transfer in operator in GB
//   - CU_NP: returns sent/received network transfer peering in GB
//   - CU_I: returns number of public IPs
func (a Account) GetConsumedCloudUnitsByType(ctx context.Context, req GetConsumedCloudUnitsByTypeRequest) (float64, error) {
	err := req.validate()
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
