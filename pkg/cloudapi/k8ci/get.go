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
		return errors.New("field K8CIID can not be empty or equal to 0")
	}

	return nil
}

// Get gets details of the specified K8 catalog item
func (k K8CI) Get(ctx context.Context, req GetRequest) (*RecordK8CI, error) {
	if err := req.validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/k8ci/get"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8CI{}
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
