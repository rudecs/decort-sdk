package disks

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListUnattachedRequest struct {
	AccountId uint64 `url:"accountId"`
}

func (d Disks) ListUnattached(ctx context.Context, req ListUnattachedRequest, options ...opts.DecortOpts) (DiskList, error) {
	url := "/disks/listUnattached"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := d.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	diskList := DiskList{}

	err = json.Unmarshal(res, &diskList)
	if err != nil {
		return nil, err
	}

	return diskList, nil

}
