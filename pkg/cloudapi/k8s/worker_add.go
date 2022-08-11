package k8s

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type WorkerAddRequest struct {
	K8SId          uint64 `url:"k8sId"`
	WorkersGroupId uint64 `url:"workersGroupId"`
	Num            uint   `url:"num"`
}

func (krq WorkerAddRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}

	if krq.WorkersGroupId == 0 {
		return errors.New("validation-error: field WorkersGroupId can not be empty or equal to 0")
	}

	if krq.Num == 0 {
		return errors.New("validation-error: field Num can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) WorkerAdd(ctx context.Context, req WorkerAddRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/k8s/workerAdd"
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