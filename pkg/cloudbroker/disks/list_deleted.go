package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list deleted disks
type ListDeletedRequest struct {
	// ID of the account the disks belong to
	// Required: false
	AccountID uint64 `url:"accountId,omitempty"`

	// Type of the disks
	// Required: false
	Type string `url:"type,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// ListDeleted gets list the deleted disks based on filter
func (d Disks) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListDeletedDisks, error) {
	url := "/cloudbroker/disks/listDeleted"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDeletedDisks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
