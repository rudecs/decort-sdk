package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete snapshot
type SnapshotDeleteRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// Label of the snapshot
	// Required: true
	Label string `url:"label" json:"label"`
}

func (bsrq SnapshotDeleteRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.Label == "" {
		return errors.New("field Label can not be empty")
	}

	return nil
}

// SnapshotDelete delete snapshot of the Basic Service
func (b BService) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/snapshotDelete"

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
