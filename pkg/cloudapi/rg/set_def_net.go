package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for set default network
type SetDefNetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Network type
	// Should be one of:
	//	- "PUBLIC"
	//	- "PRIVATE"
	// Required: true
	NetType string `url:"netType"`

	// Network ID
	// Required: false
	NetID uint64 `url:"netId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq SetDefNetRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}
	if !validators.StringInSlice(rgrq.NetType, []string{"PUBLIC", "PRIVATE"}) {
		return errors.New("field NetType can only be one of 'PUBLIC' or 'PRIVATE'")
	}

	return nil
}

// SetDefNet sets default network for attach associated virtual machines
func (r RG) SetDefNet(ctx context.Context, req SetDefNetRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/rg/setDefNet"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
