package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for calculate the currently consumed units for all cloudspaces and resource groups in the account
type GetConsumedAccountUnitsRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`
}

func (arq GetConsumedAccountUnitsRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// GetConsumedAccountUnits calculates the currently consumed units for all cloudspaces and resource groups in the account.
// Calculated cloud units are returned in a dict which includes:
//   - CU_M: consumed memory in MB
//   - CU_C: number of cpu cores
//   - CU_D: consumed vdisk storage in GB
//   - CU_I: number of public IPs
func (a Account) GetConsumedAccountUnits(ctx context.Context, req GetConsumedAccountUnitsRequest) (*ResourceLimits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/getConsumedAccountUnits"

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
