package disks

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	AccountId uint64 `url:"accountId,omitempty"`
	Type      string `url:"type,omitempty"`
	Page      uint64 `url:"page,omitempty"`
	Size      uint64 `url:"size,omitempty"`
}

func (d Disks) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (DiskList, error) {
	url := "/disks/list"
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

func (d Disks) ListDeleted(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (DiskList, error) {
	url := "/disks/listDeleted"
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
