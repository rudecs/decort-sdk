package k8s

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type WorkersGroupGetByNameRequest struct {
	K8SId     uint64 `url:"k8sId"`
	GroupName string `url:"groupName "`
}

func (krq WorkersGroupGetByNameRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}

	if krq.GroupName == "" {
		return errors.New("validation-error: field WorkersGroupId can not be empty")
	}

	return nil
}

func (k8s K8S) WorkersGroupGetByName(ctx context.Context, req WorkersGroupGetByNameRequest, options ...opts.DecortOpts) (*K8SGroup, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/k8s/workersGroupGetByName"
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
		return nil, err
	}

	group := &K8SGroup{}

	err = json.Unmarshal(res, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}
