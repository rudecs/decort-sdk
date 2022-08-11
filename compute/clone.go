package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CloneRequest struct {
	ComputeId         uint64 `url:"computeId"`
	Name              string `url:"name"`
	SnapshotTimestamp uint64 `url:"snapshotTimestamp"`
	SnapshotName      string `url:"snapshotName"`
}

func (crq CloneRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}
	if crq.Name == "" {
		return errors.New("validation-error: field Name can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Clone(ctx context.Context, req CloneRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/compute/clone"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
