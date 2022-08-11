package flipgroup

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateRequest struct {
	AccountID   uint64 `url:"accountId"`
	Name        string `url:"name"`
	NetType     string `url:"netType"`
	NetID       uint64 `url:"netId"`
	ClientType  string `url:"clientType"`
	IP          string `url:"ip,omitempty"`
	Description string `url:"desc,omitempty"`
}

func (frq CreateRequest) Validate() error {
	if frq.AccountID == 0 {
		return errors.New("field AccountID can not be empty or equal to 0")
	}
	if frq.NetID == 0 {
		return errors.New("field NetID can not be empty or equal to 0")
	}
	if frq.Name == "" {
		return errors.New("field Name can not be empty")
	}

	validator := validators.StringInSlice(frq.NetType, []string{"EXTNET", "VINS"})
	if !validator {
		return errors.New("field Name can be only EXTNET or VINS")
	}
	validator = validators.StringInSlice(frq.ClientType, []string{"compute", "node"})
	if !validator {
		return errors.New("field Name can be only compute or node")
	}

	return nil
}

func (f FlipGroup) Create(ctx context.Context, req CreateRequest, options ...opts.DecortOpts) (*FlipGroupRecord, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/flipgroup/create"
	fgRaw, err := f.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	fg := &FlipGroupRecord{}
	if err := json.Unmarshal(fgRaw, fg); err != nil {
		return nil, err
	}

	return fg, nil
}
