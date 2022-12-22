package k8ci

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about K8CI
type GetRequest struct {
	// ID of the K8 catalog item to get
	// Required: true
	K8CIID uint64 `url:"k8ciId"`
}

func (krq GetRequest) validate() error {
	if krq.K8CIID == 0 {
		return errors.New("validation-error: field K8CIID must be set")
	}

	return nil
}

// Get gets details of the specified K8 catalog item
func (k K8CI) Get(ctx context.Context, req GetRequest) (*RecordK8CI, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/k8ci/get"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	item := RecordK8CI{}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
