package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SnapshotUsageRequest struct {
	ComputeId uint64 `url:"computeId"`
	Label     string `url:"label,omitempty"`
}

func (crq SnapshotUsageRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) SnapshotUsage(ctx context.Context, req SnapshotUsageRequest, options ...opts.DecortOpts) (SnapshotUsageList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/snapshotUsage"
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

	snapshotUsage := SnapshotUsageList{}

	err = json.Unmarshal(res, &snapshotUsage)
	if err != nil {
		return nil, err
	}

	return snapshotUsage, nil
}
