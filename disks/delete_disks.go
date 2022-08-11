package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DisksDeleteRequest struct {
	DisksIds    []uint64 `url:"diskIds"`
	Reason      string   `url:"reason"`
	Permanently bool     `url:"permanently"`
}

func (drq DisksDeleteRequest) Validate() error {
	if len(drq.DisksIds) == 0 {
		return errors.New("validation-error: field DisksIds must include one or more disks ids")
	}

	return nil
}

func (d Disks) DeleteDisks(ctx context.Context, req DisksDeleteRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/disks/deleteDisks"
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
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil

}
