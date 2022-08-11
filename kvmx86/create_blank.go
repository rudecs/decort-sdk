package kvmx86

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateBlankRequest struct {
	RGID        uint64 `url:"rgId"`
	Name        string `url:"name"`
	CPU         uint64 `url:"cpu"`
	RAM         uint64 `url:"ram"`
	BootDisk    uint64 `url:"bootDisk"`
	SepID       uint64 `url:"sepId"`
	Pool        string `url:"pool"`
	NetType     string `url:"netType,omitempty"`
	NetID       uint64 `url:"netId,omitempty"`
	IPAddr      string `url:"ipAddr,omitempty"`
	Description string `url:"desc,omitempty"`
}

func (krq CreateBlankRequest) Validate() error {
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
	if krq.BootDisk == 0 {
		return errors.New("validation-error: field BootDisk can not be empty or equal to 0")
	}
	if krq.SepID == 0 {
		return errors.New("validation-error: field SepID can not be empty or equal to 0")
	}
	if krq.Pool == "" {
		return errors.New("validation-error: field Pool can not be empty")
	}

	return nil
}

func (k KVMX86) CreateBlank(ctx context.Context, req CreateBlankRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/kvmx86/createBlank"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := k.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil

}
