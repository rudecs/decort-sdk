package k8s

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type WorkersGroupAddRequest struct {
	K8SId       uint64   `url:"k8sId"`
	Name        string   `url:"name"`
	Labels      []string `url:"labels"`
	Taints      []string `url:"taints"`
	Annotations []string `url:"annotations"`
	WorkerNum   uint     `url:"workerNum"`
	WorkerCpu   uint     `url:"workerCpu"`
	WorkerRam   uint     `url:"workerRam"`
	WorkerDisk  uint     `url:"workerDisk"`
}

func (krq WorkersGroupAddRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}

	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if krq.WorkerNum == 0 {
		return errors.New("validation-error: field WorkerNum can not be empty or equal to 0")
	}

	if krq.WorkerCpu == 0 {
		return errors.New("validation-error: field WorkerCpu can not be empty or equal to 0")
	}

	if krq.WorkerRam < 1024 {
		return errors.New("validation-error: field WorkerRam must be greater or equal 1024")
	}

	return nil
}

func (k8s K8S) WorkersGroupAdd(ctx context.Context, req WorkersGroupAddRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/k8s/workersGroupAdd"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := k8s.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}
	return result, nil
}
