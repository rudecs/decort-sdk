package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete master from group
type DeleteMasterFromGroupRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId"`

	// ID of the masters compute group
	// Required: true
	MasterGroupID uint64 `url:"masterGroupId" json:"masterGroupId"`

	// List of Compute IDs of master nodes to delete
	// Required: true
	MasterIDs []string `url:"masterIds" json:"masterIds"`
}

func (krq DeleteMasterFromGroupRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.MasterGroupID == 0 {
		return errors.New("validation-error: field MasterGroupID must be set")
	}
	if len(krq.MasterIDs) == 0 {
		return errors.New("validation-error: field MasterIDs must be set")
	}

	return nil
}

// DeleteMasterFromGroup deletes compute from masters group in selected kubernetes cluster
func (k K8S) DeleteMasterFromGroup(ctx context.Context, req DeleteMasterFromGroupRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/deleteMasterFromGroup"

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
