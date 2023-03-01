package k8s

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for get node annotations
type GetNodeAnnotationsRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// Node ID
	// Required: true
	NodeID uint64 `url:"nodeId" json:"nodeId"`
}

func (krq GetNodeAnnotationsRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.NodeID == 0 {
		return errors.New("validation-error: field NodeID can not be empty or equal to 0")
	}

	return nil
}

// GetNodeAnnotations gets kubernetes cluster worker node annotations
func (k K8S) GetNodeAnnotations(ctx context.Context, req GetNodeAnnotationsRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/k8s/getNodeAnnotations"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
