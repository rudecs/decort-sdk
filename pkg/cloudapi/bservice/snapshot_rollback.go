package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for rollback snapshot
type SnapshotRollbackRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId"`

	// Label of the snapshot
	// Required: true
	Label string `url:"label"`
}

func (bsrq SnapshotRollbackRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.Label == "" {
		return errors.New("field Label can not be empty")
	}

	return nil
}

// SnapshotRollback rollback snapshot of the Basic Service
func (b BService) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/snapshotRollback"

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
