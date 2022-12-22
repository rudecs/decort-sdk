package tasks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of tasks
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty"`
}

// List gets list user API tasks with status PROCESSING
func (t Tasks) List(ctx context.Context, req ListRequest) (ListTasks, error) {
	url := "/cloudapi/tasks/list"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListTasks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil

}
