package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update kubernetes cluster
type UpdateRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// New name to set.
	// If empty string is passed, name is not updated
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// New description to set.
	// If empty string is passed, description is not updated
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

func (krq UpdateRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}

	return nil
}

// Update updates name or description of kubernetes cluster
func (k K8S) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/update"

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
