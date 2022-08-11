package k8s

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	K8SId uint64 `url:"k8sId"`
}

func (krq GetRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*K8SRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/k8s/get"
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

	k8sInfo := &K8SRecord{}

	err = json.Unmarshal(res, k8sInfo)
	if err != nil {
		return nil, err
	}

	return k8sInfo, nil
}
