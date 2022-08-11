package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SnapshotListRequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq SnapshotListRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) SnapshotList(ctx context.Context, req SnapshotListRequest, options ...opts.DecortOpts) (SnapshotList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/snapshotList"
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
		return nil, err
	}

	snapshotList := SnapshotList{}

	err = json.Unmarshal(res, &snapshotList)
	if err != nil {
		return nil, err
	}

	return snapshotList, nil
}
