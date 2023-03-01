package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get background API task status and result
type GetRequest struct {
	// ID of audit GUID
	// Required: true
	AuditID string `url:"auditId" json:"auditId"`
}

func (trq GetRequest) validate() error {
	if trq.AuditID == "" {
		return errors.New("validation-error: field AuditID must be set")
	}

	return nil
}

// Get gets background API task status and result
func (t Tasks) Get(ctx context.Context, req GetRequest) (*RecordTask, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/tasks/get"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	item := RecordTask{}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
