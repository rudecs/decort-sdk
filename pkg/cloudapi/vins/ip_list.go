package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for DHCP IP
type IPListRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`
}

func (vrq IPListRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

// IPList show DHCP IP reservations on VINS
func (v VINS) IPList(ctx context.Context, req IPListRequest) (ListIPs, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/vins/ipList"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListIPs{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
