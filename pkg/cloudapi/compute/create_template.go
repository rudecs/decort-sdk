package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type CreateTemplateRequest struct {
	ComputeID uint64 `url:"computeId"`
	Name      string `url:"name"`
	Async     bool   `url:"async"`
}

func (crq CreateTemplateRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	return nil
}

func (c Compute) CreateTemplate(ctx context.Context, req CreateTemplateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	req.Async = false

	url := "/cloudapi/compute/createTemplate"

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

func (c Compute) CreateTemplateAsync(ctx context.Context, req CreateTemplateRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	req.Async = true

	url := "/cloudapi/compute/createTemplate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
