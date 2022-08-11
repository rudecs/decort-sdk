package k8s

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	IncludeDeleted bool   `json:"includedeleted"`
	Page           uint64 `json:"page"`
	Size           uint64 `json:"size"`
}

func (k8s K8S) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (K8SList, error) {

	url := "/k8s/list"
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

	k8sList := K8SList{}

	err = json.Unmarshal(res, &k8sList)
	if err != nil {
		return nil, err
	}

	return k8sList, nil
}
