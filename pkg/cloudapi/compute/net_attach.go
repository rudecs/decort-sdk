package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type NetAttachRequest struct {
	ComputeID uint64 `url:"computeId"`
	NetType   string `url:"netType"`
	NetID     uint64 `url:"netId"`
	IPAddr    string `url:"ipAddr,omitempty"`
}

func (crq NetAttachRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
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

func (c Compute) NetAttach(ctx context.Context, req NetAttachRequest) (*NetAttach, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/netAttach"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
