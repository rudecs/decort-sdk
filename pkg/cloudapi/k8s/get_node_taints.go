package k8s

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for get node taints
type GetNodeTaintsRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// Node ID
	// Required: false
	NodeID uint64 `url:"nodeId" json:"nodeId"`
}

func (krq GetNodeTaintsRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.NodeID == 0 {
		return errors.New("validation-error: field NodeID can not be empty or equal to 0")
	}

	return nil
}

// GetNodeTaints gets kubernetes cluster worker node taints
func (k8s K8S) GetNodeTaints(ctx context.Context, req GetNodeTaintsRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/k8s/getNodeTaints"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
