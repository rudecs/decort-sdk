package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list types of disks
type ListTypesRequest struct {
	// Show detailed disk types by seps
	// Required: false
	Detailed bool `url:"detailed"`
}

// ListTypes gets list defined disk types
func (d Disks) ListTypes(ctx context.Context, req ListTypesRequest) ([]interface{}, error) {
	url := "/cloudapi/disks/listTypes"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]interface{}, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
