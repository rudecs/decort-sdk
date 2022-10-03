package kvmppc

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	RGID        uint64 `url:"rgId"`
	Name        string `url:"name"`
	CPU         uint64 `url:"cpu"`
	RAM         uint64 `url:"ram"`
	ImageID     uint64 `url:"imageId"`
	BootDisk    uint64 `url:"bootDisk,omitempty"`
	SepID       uint64 `url:"sepId,omitempty"`
	Pool        string `url:"pool,omitempty"`
	NetType     string `url:"netType,omitempty"`
	NetID       uint64 `url:"netId,omitempty"`
	IPAddr      string `url:"ipAddr,omitempty"`
	Userdata    string `url:"userdata,omitempty"`
	Description string `url:"desc,omitempty"`
	Start       bool   `url:"start,omitempty"`
	IS          string `url:"IS,omitempty"`
	IPAType     string `url:"ipaType,omitempty"`
}

func (krq CreateRequest) Validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU can not be empty or equal to 0")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM can not be empty or equal to 0")
	}
	if krq.ImageID == 0 {
		return errors.New("validation-error: field ImageID can not be empty or equal to 0")
	}

	return nil
}

func (k KVMPPC) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/kvmppc/create"
	prefix := "/cloudapi"

	url = prefix + url
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
