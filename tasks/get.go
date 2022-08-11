package tasks

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	AuditId string `url:"auditId"`
}

func (trq GetRequest) Validate() error {
	if trq.AuditId == "" {
		return errors.New("validation-error: field AuditId can not be empty")
	}

	return nil
}

func (t Tasks) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*AsyncTask, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/tasks/get"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	taskRaw, err := t.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	task := &AsyncTask{}
	err = json.Unmarshal(taskRaw, task)
	if err != nil {
		return nil, err
	}

	return task, nil

}
