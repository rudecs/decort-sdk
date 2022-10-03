package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListUnattachedRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (d Disks) ListUnattached(ctx context.Context, req ListUnattachedRequest) (DiskList, error) {
	url := "/cloudapi/disks/listUnattached"

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
