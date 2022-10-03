package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateInAccountRequest struct {
	Name               string `url:"name"`
	AccountID          uint64 `url:"accountId"`
	GID                uint64 `url:"gid,omitempty"`
	IPCidr             string `url:"ipcidr,omitempty"`
	Description        string `url:"desc,omitempty"`
	PreReservationsNum uint   `url:"preReservationsNum,omitempty"`
}

func (vrq CreateInAccountRequest) Validate() error {
	if vrq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if vrq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) CreateInAccount(ctx context.Context, req CreateInAccountRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/vins/createInAccount"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
