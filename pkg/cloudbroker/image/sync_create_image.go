package image

import (
	"context"
	"errors"
	"github.com/rudecs/decort-sdk/internal/validators"
	"net/http"
	"strconv"
)

type SyncCreateRequest struct {
	Name         string   `url:"name"`
	URL          string   `url:"url"`
	GID          uint64   `url:"gid"`
	BootType     string   `url:"boottype"`
	ImageType    string   `url:"imagetype"`
	HotResize    bool     `url:"hotresize,omitempty"`
	Username     string   `url:"username,omitempty"`
	Password     string   `url:"password,omitempty"`
	AccountID    uint64   `url:"accountId,omitempty"`
	UsernameDL   string   `url:"usernameDL,omitempty"`
	PasswordDL   string   `url:"passwordDL,omitempty"`
	SepID        uint64   `url:"sepId,omitempty"`
	PoolName     string   `url:"poolName,omitempty"`
	Architecture string   `url:"architecture,omitempty"`
	Drivers      []string `url:"drivers"`
	Bootable     bool     `url:"bootable,omitempty"`
}

func (irq SyncCreateRequest) Validate() error {
	if irq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if irq.URL == "" {
		return errors.New("validation-error: field URL must be set")
	}
	if irq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if irq.BootType == "" {
		return errors.New("validation-error: field BootType must be set")
	}
	if irq.ImageType == "" {
		return errors.New("validation-error: field ImageType must be set")
	}

	validate := validators.StringInSlice(irq.BootType, []string{"bios", "uefi"})
	if !validate {
		return errors.New("validation-error: field BootType can be bios or uefi")
	}
	validate = validators.StringInSlice(irq.ImageType, []string{"windows", "linux", "other"})
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

func (i Image) SyncCreate(ctx context.Context, req SyncCreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/image/syncCreateImage"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}