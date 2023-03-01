package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Request struct for create template
type CreateTemplateRequest struct {
	// ID of the compute to create template from
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Name to assign to the template being created
	// Required: true
	Name string `url:"name" json:"name"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`

	// Async API call
	// For async call use CreateTemplateAsync
	// For sync call use CreateTemplate
	// Required: true
	async bool `url:"async"`
}

func (crq CreateTemplateRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// CreateTemplateAsync create template from compute instance
func (c Compute) CreateTemplateAsync(ctx context.Context, req CreateTemplateRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	req.async = true

	url := "/cloudbroker/compute/createTemplate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}

// CreateTemplate create template from compute instance
func (c Compute) CreateTemplate(ctx context.Context, req CreateTemplateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	req.async = false

	url := "/cloudbroker/compute/createTemplate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
