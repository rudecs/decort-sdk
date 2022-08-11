package disks

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListTypesRequest struct {
	Detailed bool `url:"detailed"`
}

func (d Disks) ListTypes(ctx context.Context, req ListTypesRequest, options ...opts.DecortOpts) ([]interface{}, error) {
	url := "/disks/listTypes"
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

	typesList := make([]interface{}, 0)

	err = json.Unmarshal(res, &typesList)
	if err != nil {
		return nil, err
	}

	return typesList, nil

}
