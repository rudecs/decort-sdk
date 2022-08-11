package bservice

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SnapshotListRequest struct {
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq SnapshotListRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) SnapshotList(ctx context.Context, req SnapshotListRequest, options ...opts.DecortOpts) ([]Snapshot, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/bservice/snapshotList"
	snapshotListRaw, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	snapshotList := []Snapshot{}
	if err := json.Unmarshal(snapshotListRaw, &snapshotList); err != nil {
		return nil, err
	}

	return snapshotList, nil
}
