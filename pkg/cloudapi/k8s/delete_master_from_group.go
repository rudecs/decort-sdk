package k8s

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteMasterFromGroupRequest struct {
	K8SId         uint64   `url:"k8sId"`
	MasterGroupId uint64   `url:"masterGroupId"`
	MasterIds     []string `url:"masterIds"`
}

func (krq DeleteMasterFromGroupRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}
	if krq.MasterGroupId == 0 {
		return errors.New("validation-error: field MasterGroupId can not be empty or equal to 0")
	}

	if len(krq.MasterIds) == 0 {
		return errors.New("validation-error: field MasterIds can not be empty")
	}

	return nil
}

func (k8s K8S) DeleteMasterFromGroup(ctx context.Context, req DeleteMasterFromGroupRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/k8s/deleteMasterFromGroup"
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
