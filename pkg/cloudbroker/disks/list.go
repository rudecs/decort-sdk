package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	AccountID uint64 `url:"accountId,omitempty"`
	Type      string `url:"type,omitempty"`
	Page      uint64 `url:"page,omitempty"`
	Size      uint64 `url:"size,omitempty"`
}

func (d Disks) List(ctx context.Context, req ListRequest) (ListDisks, error) {
	url := "/cloudbroker/disks/list"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	listDisks := ListDisks{}

	err = json.Unmarshal(res, &listDisks)
	if err != nil {
		return nil, err
	}

	return listDisks, nil
}
