package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get infromation about task
type GetRequest struct {
	// ID of audit
	// Required: true
	AuditID string `url:"auditId"`
}

func (trq GetRequest) validate() error {
	if trq.AuditID == "" {
		return errors.New("validation-error: field AuditID can not be empty")
	}

	return nil
}

// Get gets background API task status and result
func (t Tasks) Get(ctx context.Context, req GetRequest) (*RecordAsyncTask, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/tasks/get"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordAsyncTask{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil

}
