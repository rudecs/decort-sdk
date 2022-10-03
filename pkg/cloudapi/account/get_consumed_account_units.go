package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetConsumedAccountUnitsRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (arq GetConsumedAccountUnitsRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (a Account) GetConsumedAccountUnits(ctx context.Context, req GetConsumedAccountUnitsRequest) (*ResourceLimits, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/account/getConsumedAccountUnits"

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
