package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for calculate the reserved units for all cloudspaces and resource groups in the account
type GetReservedAccountUnitsRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`
}

func (arq GetReservedAccountUnitsRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// GetReservedAccountUnits calculates the reserved units for all cloudspaces and resource groups in the account.
// Calculated cloud units are returned in a dict which includes:
//
//   - CU_M: reserved memory in MB
//   - CU_C: number of cpu cores
//   - CU_D: reserved vdisk storage in GB
//   - CU_I: number of public IPs
func (a Account) GetReservedAccountUnits(ctx context.Context, req GetReservedAccountUnitsRequest) (*ResourceLimits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/getReservedAccountUnits"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := ResourceLimits{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
