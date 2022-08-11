package k8s

import (
	"context"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetConfigRequest struct {
	K8SId uint64 `url:"k8sId"`
}

func (krq GetConfigRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) GetConfig(ctx context.Context, req GetConfigRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/k8s/getConfig"
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
		return "", err
	}

	return string(res), nil
}
