package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteMasterFromGroupRequest struct {
	K8SID         uint64   `url:"k8sId"`
	MasterGroupID uint64   `url:"masterGroupId"`
	MasterIDs     []string `url:"masterIds"`
}

func (krq DeleteMasterFromGroupRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.MasterGroupID == 0 {
		return errors.New("validation-error: field MasterGroupID can not be empty or equal to 0")
	}

	if len(krq.MasterIDs) == 0 {
		return errors.New("validation-error: field MasterIDs can not be empty")
	}

	return nil
}

func (k8s K8S) DeleteMasterFromGroup(ctx context.Context, req DeleteMasterFromGroupRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/deleteMasterFromGroup"

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
