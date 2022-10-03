package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateInRGRequest struct {
	Name               string `url:"name"`
	RGID               uint64 `url:"rgId"`
	IPCidr             string `url:"ipcidr,omitempty"`
	ExtNetID           uint64 `url:"extNetId,omitempty"`
	ExtIP              string `url:"extIp,omitempty"`
	Description        string `url:"desc,omitempty"`
	PreReservationsNum uint   `url:"preReservationsNum,omitempty"`
}

func (vrq CreateInRGRequest) Validate() error {
	if vrq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if vrq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) CreateInRG(ctx context.Context, req CreateInRGRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/vins/createInRG"

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
