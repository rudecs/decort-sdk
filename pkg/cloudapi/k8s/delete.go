package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete kubernetes cluster
type DeleteRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// True if cluster is destroyed permanently.
	// Otherwise it can be restored from Recycle Bin
	// Required: true
	Permanently bool `url:"permanently" json:"permanently"`
}

func (krq DeleteRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	return nil
}

// Delete deletes kubernetes cluster
func (k8s K8S) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/delete"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
