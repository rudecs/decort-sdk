package kvmppc

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	RGID     uint64 `url:"rgId"`
	Name     string `url:"name"`
	CPU      uint64 `url:"cpu"`
	RAM      uint64 `url:"ram"`
	ImageID  uint64 `url:"imageId"`
	BootDisk uint64 `url:"bootDisk,omitempty"`
	SepID    uint64 `url:"sepId,omitempty"`
	Pool     string `url:"pool,omitempty"`
	NetType  string `url:"netType,omitempty"`
	NetID    uint64 `url:"netId,omitempty"`
	IPAddr   string `url:"ipAddr,omitempty"`
	UserData string `url:"userdata,omitempty"`
	Desc     string `url:"desc,omitempty"`
	Start    bool   `url:"start,omitempty"`
	StackID  uint64 `url:"stackId,omitempty"`
	IS       string `url:"IS,omitempty"`
	IpaType  string `url:"ipaType,omitempty"`
	Reason   string `url:"reason,omitempty"`
}

func (krq CreateRequest) Validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU must be set")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM must be set")
	}
	if krq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (k KVMPPC) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/kvmppc/create"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
