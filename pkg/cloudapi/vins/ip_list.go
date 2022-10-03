package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type IPListRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq IPListRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) IPList(ctx context.Context, req IPListRequest) (IPList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/vins/ipList"

	ipListRaw, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	ipList := IPList{}
	err = json.Unmarshal(ipListRaw, &ipList)
	if err != nil {
		return nil, err
	}

	return ipList, nil

}
