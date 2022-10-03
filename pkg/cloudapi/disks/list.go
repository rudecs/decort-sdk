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

func (d Disks) List(ctx context.Context, req ListRequest) (DiskList, error) {
	url := "/cloudapi/disks/list"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	diskList := DiskList{}

	err = json.Unmarshal(res, &diskList)
	if err != nil {
		return nil, err
	}

	return diskList, nil

}

func (d Disks) ListDeleted(ctx context.Context, req ListRequest) (DiskList, error) {
	url := "/disks/listDeleted"
	prefix := "/cloudapi"

	url = prefix + url
	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	diskList := DiskList{}

	err = json.Unmarshal(res, &diskList)
	if err != nil {
		return nil, err
	}

	return diskList, nil

}
