package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AffinityRelationsRequest struct {
	ComputeId     uint64 `url:"computeId"`
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityRelationsRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) AffinityRelations(ctx context.Context, req AffinityRelationsRequest, options ...opts.DecortOpts) (*AffinityRelations, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/affinityRelations"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	relations := &AffinityRelations{}

	err = json.Unmarshal([]byte(res), relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}
