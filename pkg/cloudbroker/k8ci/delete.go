package k8ci

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete K8CI
type DeleteRequest struct {
	// K8CI ID
	// Required: true
	K8CIID uint64 `url:"k8ciId"`

	// Delete permanently or not
	// Required: false
	Permanently bool `url:"permanently,omitempty"`
}

func (krq DeleteRequest) validate() error {
	if krq.K8CIID == 0 {
		return errors.New("validation-error: field K8CIID must be set")
	}

	return nil
}

// Delete deletes K8CI by ID
func (k K8CI) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8ci/delete"

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
