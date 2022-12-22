package flipgroup

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create FLIPGroup
type CreateRequest struct {
	// Account ID
	// Required: true
	AccountID uint64 `url:"accountId"`

	// FLIPGroup name
	// Required: true
	Name string `url:"name"`

	// Network type
	// Should be one of:
	//	- EXTNET
	//	- VINS
	// Required: true
	NetType string `url:"netType"`

	// ID of external network or VINS
	// Required: true
	NetID uint64 `url:"netId"`

	// Type of client
	//	- 'compute'
	//	- 'vins' (will be later)
	// Required: true
	ClientType string `url:"clientType"`

	// IP address to associate with this group. If empty, the platform will autoselect IP address
	// Required: false
	IP string `url:"ip,omitempty"`

	// Text description of this FLIPGorup instance
	// Required: false
	Description string `url:"desc,omitempty"`
}

func (frq CreateRequest) validate() error {
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

// Create method will create a new FLIPGorup in the specified Account
func (f FLIPGroup) Create(ctx context.Context, req CreateRequest) (*RecordFLIPGroup, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/flipgroup/create"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordFLIPGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
