package vins

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type IPListRequest struct {
	VinsId uint64 `url:"vinsId"`
}

func (vrq IPListRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) IPList(ctx context.Context, req IPListRequest, options ...opts.DecortOpts) (IPList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/vins/ipList"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	ipListRaw, err := v.client.DecortApiCall(ctx, typed.POST, url, req)
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
