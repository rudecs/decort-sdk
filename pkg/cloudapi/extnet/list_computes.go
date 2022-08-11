package extnet

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListComputesRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (erq ListComputesRequest) Validate() error {
	if erq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}

	return nil
}

func (e Extnet) ListComputes(ctx context.Context, req ListComputesRequest, options ...opts.DecortOpts) (ExtnetComputesList, error) {
	url := "/extnet/listComputes"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	extnetComputesListRaw, err := e.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	extnetComputesList := ExtnetComputesList{}
	err = json.Unmarshal(extnetComputesListRaw, &extnetComputesList)
	if err != nil {
		return nil, err
	}

	return extnetComputesList, nil

}
