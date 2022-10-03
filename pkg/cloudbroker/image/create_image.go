package image

import (
	"context"
	"errors"
	"github.com/rudecs/decort-sdk/internal/validators"
	"net/http"
	"strconv"
)

type CreateRequest struct {
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
	Drivers      []string `url:"drivers,omitempty"`
	Bootable     bool     `url:"bootable,omitempty"`
}

func (irq CreateRequest) Validate() error {
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

	return nil
}

func (i Image) CreateImage(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/image/createImage"

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
