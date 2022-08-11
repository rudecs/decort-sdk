package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type NetAttachRequest struct {
	ComputeId uint64 `url:"computeId"`
	NetType   string `url:"netType"`
	NetID     uint64 `url:"netId"`
	IPAddr    string `url:"ipAddr,omitempty"`
}

func (crq NetAttachRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}
	if crq.NetType == "" {
		return errors.New("validation-error: field NetType can not be empty")
	}
	validator := validators.StringInSlice(crq.NetType, []string{"EXTNET", "VINS"})
	if !validator {
		return errors.New("validation-error: field NetType can be only EXTNET or VINS")
	}

	if crq.NetID == 0 {
		return errors.New("validation-error: field NetID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) NetAttach(ctx context.Context, req NetAttachRequest, options ...opts.DecortOpts) (*NetAttach, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/netAttach"
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

	netAttach := &NetAttach{}
	err = json.Unmarshal(res, netAttach)
	if err != nil {
		return nil, err
	}

	return netAttach, nil
}
