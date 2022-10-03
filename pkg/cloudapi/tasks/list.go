package tasks

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (t Tasks) List(ctx context.Context, req ListRequest) (TasksList, error) {
	url := "/tasks/list"
	prefix := "/cloudapi"

	url = prefix + url
	taskListRaw, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	taskList := TasksList{}
	err = json.Unmarshal(taskListRaw, &taskList)
	if err != nil {
		return nil, err
	}

	return taskList, nil

}
