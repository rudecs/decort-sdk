package vins

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ExtNetListRequest struct {
	VinsId uint64 `url:"vinsId"`
}

func (vrq ExtNetListRequest) Validate() error {
	if vrq.VinsId == 0 {
		return errors.New("validation-error: field VinsId can not be empty or equal to 0")
	}

	return nil
}

func (v Vins) ExtNetList(ctx context.Context, req ExtNetListRequest, options ...opts.DecortOpts) (ExtnetList, error) {
	url := "/vins/extNetList"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	extnetListRaw, err := v.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	extnetList := ExtnetList{}
	err = json.Unmarshal(extnetListRaw, &extnetList)
	if err != nil {
		return nil, err
	}

	return extnetList, nil

}
