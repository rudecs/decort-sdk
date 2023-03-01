package tasks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list audits
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list user API task with status PROCESSING
func (t Tasks) List(ctx context.Context, req ListRequest) (ListTasks, error) {
	url := "/cloudbroker/tasks/list"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	item := ListTasks{}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
