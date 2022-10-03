package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	AuditID string `url:"auditId"`
}

func (trq GetRequest) Validate() error {
	if trq.AuditID == "" {
		return errors.New("validation-error: field AuditID can not be empty")
	}

	return nil
}

func (t Tasks) Get(ctx context.Context, req GetRequest) (*AsyncTask, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/tasks/get"
	prefix := "/cloudapi"

	url = prefix + url
	taskRaw, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
