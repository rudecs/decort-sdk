package image

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateRequest struct {
	Name         string   `url:"name"`
	URL          string   `url:"url"`
	GID          uint64   `url:"gid"`
	BootType     string   `url:"boottype"`
	ImageType    string   `url:"imagetype"`
	Hotresize    bool     `url:"hotresize,omitempty"`
	Username     string   `url:"username,omitempty"`
	Password     string   `url:"password,omitempty"`
	AccountId    uint64   `url:"accountId,omitempty"`
	UsernameDL   string   `url:"usernameDL,omitempty"`
	PasswordDL   string   `url:"passwordDL,omitempty"`
	SepId        uint64   `url:"sepId,omitempty"`
	Pool         string   `url:"poolName,omitempty"`
	Architecture string   `url:"architecture,omitempty"`
	Drivers      []string `url:"drivers"`
}

func (irq CreateRequest) Validate() error {
	if irq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if irq.URL == "" {
		return errors.New("validation-error: field URL can not be empty")
	}

	if irq.GID == 0 {
		return errors.New("validation-error: field GID can not be empty or equal to 0")
	}
	if irq.BootType == "" {
		return errors.New("validation-error: field BootType can not be empty")
	}

	validate := validators.StringInSlice(irq.BootType, []string{"bios", "uefi"})
	if !validate {
		return errors.New("validation-error: field BootType can be bios or uefi")
	}
	if irq.ImageType == "" {
		return errors.New("validation-error: field ImageType can not be empty")
	}

	validate = validators.StringInSlice(irq.ImageType, []string{"bios", "uefi"})
	if !validate {
		return errors.New("validation-error: field ImageType can be windows, linux or other")
	}

	if len(irq.Drivers) == 0 || len(irq.Drivers) > 1 {
		return errors.New("validation-error: field Drivers can not be empty or have 2 or more elements")
	}

	for _, v := range irq.Drivers {
		validate := validators.StringInSlice(v, []string{"KVM_X86"})
		if !validate {
			return errors.New("validation-error: field Drivers can be KVM_X86 only")
		}
	}

	return nil
}

func (i Image) Create(ctx context.Context, req CreateRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/image/create"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := i.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}
