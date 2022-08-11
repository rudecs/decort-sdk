package disks

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SearchRequest struct {
	AccountId uint64 `url:"accountId,omitempty"`
	Name      string `url:"name,omitempty"`
	ShowAll   bool   `url:"show_all,omitempty"`
}

func (d Disks) Search(ctx context.Context, req SearchRequest, options ...opts.DecortOpts) (DiskList, error) {
	url := "/disks/search"
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
