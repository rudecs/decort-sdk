package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about group of kubernetes cluster
type FindGroupByLabelRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// List of labels to search
	// Required: true
	Labels []string `url:"labels"`

	// If true and more than one label provided, select only groups that have all provided labels.
	// If false - groups that have at least one label
	// Required: true
	Strict bool `url:"strict"`
}

func (krq FindGroupByLabelRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if len(krq.Labels) == 0 {
		return errors.New("validation-error: field Labels can not be empty")
	}

	return nil
}

// FindGroupByLabel find worker group information by one on more labels
func (k8s K8S) FindGroupByLabel(ctx context.Context, req FindGroupByLabelRequest) (ListK8SGroups, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/k8s/findGroupByLabel"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8SGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
