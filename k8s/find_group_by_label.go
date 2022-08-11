package k8s

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type FindGroupByLabelRequest struct {
	K8SId  uint64   `url:"k8sId"`
	Labels []string `url:"labels"`
	Strict bool     `url:"strict"`
}

func (krq FindGroupByLabelRequest) Validate() error {
	if krq.K8SId == 0 {
		return errors.New("validation-error: field K8SId can not be empty or equal to 0")
	}

	if len(krq.Labels) == 0 {
		return errors.New("validation-error: field Labels can not be empty")
	}

	return nil
}

func (k8s K8S) FindGroupByLabel(ctx context.Context, req FindGroupByLabelRequest, options ...opts.DecortOpts) (K8SGroupList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/k8s/findGroupByLabel"
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

	k8sGroupList := K8SGroupList{}

	err = json.Unmarshal(res, &k8sGroupList)
	if err != nil {
		return nil, err
	}

	return k8sGroupList, nil
}
