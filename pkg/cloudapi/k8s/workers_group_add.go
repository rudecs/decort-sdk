package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type WorkersGroupAddRequest struct {
	K8SID       uint64   `url:"k8sId"`
	Name        string   `url:"name"`
	Labels      []string `url:"labels"`
	Taints      []string `url:"taints"`
	Annotations []string `url:"annotations"`
	WorkerNum   uint     `url:"workerNum"`
	WorkerCPU   uint     `url:"workerCpu"`
	WorkerRam   uint     `url:"workerRam"`
	WorkerDisk  uint     `url:"workerDisk"`
}

func (krq WorkersGroupAddRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if krq.WorkerNum == 0 {
		return errors.New("validation-error: field WorkerNum can not be empty or equal to 0")
	}

	if krq.WorkerCPU == 0 {
		return errors.New("validation-error: field WorkerCPU can not be empty or equal to 0")
	}

	if krq.WorkerRam < 1024 {
		return errors.New("validation-error: field WorkerRam must be greater or equal 1024")
	}

	return nil
}

func (k8s K8S) WorkersGroupAdd(ctx context.Context, req WorkersGroupAddRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/workersGroupAdd"

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
