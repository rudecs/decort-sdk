package tasks

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (t Tasks) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (TasksList, error) {
	url := "/tasks/list"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	taskListRaw, err := t.client.DecortApiCall(ctx, typed.POST, url, req)
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
