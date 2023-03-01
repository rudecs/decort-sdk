package account

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for download the resources tracking files for an account
type GetConsumtionRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`

	// Epoch represents the start time
	// Required: true
	Start uint64 `url:"start" json:"start"`

	// Epoch represents the end time
	// Required: true
	End uint64 `url:"end" json:"end"`
}

func (arq GetConsumtionRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}
	if arq.Start == 0 {
		return errors.New("validation-error: field Start can not be empty or equal to 0")
	}
	if arq.End == 0 {
		return errors.New("validation-error: field End can not be empty or equal to 0")
	}

	return nil
}

// GetConsumtion downloads the resources tracking files for an account within a given period
func (a Account) GetConsumtion(ctx context.Context, req GetConsumtionRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/account/getConsumtion"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}

// GetConsumtionGet downloads the resources tracking files for an account within a given period
func (a Account) GetConsumtionGet(ctx context.Context, req GetConsumtionRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi//account/getConsumtion"

	res, err := a.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
