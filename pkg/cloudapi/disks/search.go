package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	AccountID uint64 `url:"accountId,omitempty"`
	Name      string `url:"name,omitempty"`
	ShowAll   bool   `url:"show_all,omitempty"`
}

func (d Disks) Search(ctx context.Context, req SearchRequest) (DiskList, error) {
	url := "/cloudapi/disks/search"

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
