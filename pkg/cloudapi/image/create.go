package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create image
type CreateRequest struct {
	// Name of the rescue disk
	// Required: true
	Name string `url:"name" json:"name"`

	// URL where to download media from
	// Required: true
	URL string `url:"url" json:"url"`

	// Grid (platform) ID where this template should be create in
	// Required: true
	GID uint64 `url:"gid" json:"gid"`

	// Boot type of image bios or UEFI
	// Required: true
	BootType string `url:"boottype" json:"boottype"`

	// Image type
	// Should be:
	//	- linux
	//	- windows
	//	- or other
	// Required: true
	ImageType string `url:"imagetype" json:"imagetype"`

	// Does this machine supports hot resize
	// Required: false
	HotResize bool `url:"hotresize,omitempty" json:"hotresize,omitempty"`

	// Optional username for the image
	// Required: false
	Username string `url:"username,omitempty" json:"username,omitempty"`

	// Optional password for the image
	// Required: false
	Password string `url:"password,omitempty" json:"password,omitempty"`

	// Account ID to make the image exclusive
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Username for upload binary media
	// Required: false
	UsernameDL string `url:"usernameDL,omitempty" json:"usernameDL,omitempty"`

	// Password for upload binary media
	// Required: false
	PasswordDL string `url:"passwordDL,omitempty" json:"passwordDL,omitempty"`

	// Storage endpoint provider ID
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool for image create
	// Required: false
	Pool string `url:"poolName,omitempty" json:"poolName,omitempty"`

	// Binary architecture of this image
	// Should be:
	//	- X86_64
	//	- PPC64_LE
	// Required: false
	Architecture string `url:"architecture,omitempty" json:"architecture,omitempty"`

	// List of types of compute suitable for image
	// Example: [ "KVM_X86" ]
	// Required: true
	Drivers []string `url:"drivers" json:"drivers"`
}

func (irq CreateRequest) validate() error {
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

// Create creates image from a media identified by URL
func (i Image) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/image/create"

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
