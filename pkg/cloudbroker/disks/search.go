package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for search
type SearchRequest struct {
	//  ID of the account to search for the Disk
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`
	// Name of the Disk to search for
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// If false, then disks having one of the statuses are not listed
	// Required: false
	ShowAll bool `url:"show_all,omitempty" json:"show_all,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// Search search disks
func (d Disks) Search(ctx context.Context, req SearchRequest) (ListDisks, error) {
	url := "/cloudbroker/disks/search"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDisks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
