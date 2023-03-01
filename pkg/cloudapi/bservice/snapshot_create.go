package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create snapshot
type SnapshotCreateRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// Label of the snapshot
	// Required: true
	Label string `url:"label" json:"label"`
}

func (bsrq SnapshotCreateRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.Label == "" {
		return errors.New("field Label can not be empty")
	}

	return nil
}

// SnapshotCreate create snapshot of the Basic Service
func (b BService) SnapshotCreate(ctx context.Context, req SnapshotCreateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/snapshotCreate"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
