package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type WorkerAddRequest struct {
	K8SID          uint64 `url:"k8sId"`
	WorkersGroupID uint64 `url:"workersGroupId"`
	Num            uint   `url:"num"`
}

func (krq WorkerAddRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID can not be empty or equal to 0")
	}

	if krq.Num == 0 {
		return errors.New("validation-error: field Num can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) WorkerAdd(ctx context.Context, req WorkerAddRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/workerAdd"

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
