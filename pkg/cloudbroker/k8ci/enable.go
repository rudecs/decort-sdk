package k8ci

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for enable K8CI
type EnableRequest struct {
	// K8CI ID
	// Required: true
	K8CIID uint64 `url:"k8ciId"`
}

func (krq EnableRequest) validate() error {
	if krq.K8CIID == 0 {
		return errors.New("validation-error: field K8CIID must be set")
	}

	return nil
}

// Enable enables K8CI
func (k K8CI) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8ci/enable"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
