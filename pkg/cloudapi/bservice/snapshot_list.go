package bservice

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list existing snapshots
type SnapshotListRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq SnapshotListRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// SnapshotList gets list existing snapshots of the Basic Service
func (b BService) SnapshotList(ctx context.Context, req SnapshotListRequest) (ListSnapshots, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/bservice/snapshotList"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSnapshots{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
