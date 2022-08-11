package locations

import (
	"context"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

func (l Locations) GetUrl(ctx context.Context, options ...opts.DecortOpts) (string, error) {
	url := "/locations/getUrl"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := l.client.DecortApiCall(ctx, typed.POST, url, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
