package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list unattached disk
type ListUnattachedRequest struct {
	// ID of the account
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListUnattached gets list of unattached disks
func (d Disks) ListUnattached(ctx context.Context, req ListUnattachedRequest) (ListUnattachedDisks, error) {
	url := "/cloudbroker/disks/listUnattached"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListUnattachedDisks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
