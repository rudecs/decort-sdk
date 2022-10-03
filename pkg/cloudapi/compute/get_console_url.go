package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type GetConsoleURLRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq GetConsoleURLRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) GetConsoleURL(ctx context.Context, req GetConsoleURLRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/compute/getConsoleUrl"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
