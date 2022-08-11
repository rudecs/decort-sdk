package locations

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (l Locations) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (LocationsList, error) {
	url := "/locations/list"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	locationsListRaw, err := l.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	locationsList := LocationsList{}
	err = json.Unmarshal(locationsListRaw, &locationsList)
	if err != nil {
		return nil, err
	}

	return locationsList, nil

}
