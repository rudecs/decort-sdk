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
	Page      uint64 `url:"page,omitempty"`
	Size      uint64 `url:"size,omitempty"`
}

func (d Disks) Search(ctx context.Context, req SearchRequest) (ListDisks, error) {
	url := "/cloudbroker/disks/search"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	diskList := ListDisks{}

	err = json.Unmarshal(res, &diskList)
	if err != nil {
		return nil, err
	}

	return diskList, nil

}
