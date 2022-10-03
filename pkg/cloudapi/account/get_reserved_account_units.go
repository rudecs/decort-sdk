package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetReservedAccountUnitsRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq GetReservedAccountUnitsRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) GetReservedAccountUnits(ctx context.Context, req GetReservedAccountUnitsRequest) (*ResourceLimits, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/getReservedAccountUnits"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
