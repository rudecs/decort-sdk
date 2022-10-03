package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq GetRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) Get(ctx context.Context, req GetRequest) (*VINSDetailed, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/vins/get"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	VINS := &VINSDetailed{}

	err = json.Unmarshal(res, VINS)
	if err != nil {
		return nil, err
	}

	return VINS, nil

}
