package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type SnapshotRollbackRequest struct {
	ServiceID uint64 `url:"serviceId"`
	Label     string `url:"label"`
}

func (bsrq SnapshotRollbackRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.Label == "" {
		return errors.New("field Label can not be empty")
	}

	return nil
}

func (b BService) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/snapshotRollback"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
