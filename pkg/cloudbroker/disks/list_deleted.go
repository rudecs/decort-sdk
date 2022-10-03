package disks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListDeletedRequest struct {
	AccountID uint64 `url:"accountId"`
	Type      string `url:"type,omitempty"`
	Page      uint64 `url:"page,omitempty"`
	Size      uint64 `url:"size,omitempty"`
}

func (drq ListDeletedRequest) Validate() error {
	if drq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

func (d Disks) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListDeletedDisks, error) {
	url := "/cloudbroker/disks/listDeleted"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	deletedDisks := ListDeletedDisks{}

	err = json.Unmarshal(res, &deletedDisks)
	if err != nil {
		return nil, err
	}

	return deletedDisks, nil
}
