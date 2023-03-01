package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create CD-ROM image
type CreateCDROMImageRequest struct {
	// Name of the rescue disk
	// Required: true
	Name string `url:"name" json:"name"`

	// URL where to download ISO from
	// Required: true
	URL string `url:"url" json:"url"`

	// Grid (platform) ID where this CD-ROM image should be create in
	// Required: true
	GID uint64 `url:"gid" json:"gid"`

	// Account ID to make the image exclusive
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Storage endpoint provider ID for place rescue CD
	// Required: false
	SEPID uint64 `url:"sep_id,omitempty" json:"sep_id,omitempty"`

	// Pool for place rescue CD
	// Required: false
	PoolName string `url:"pool_name,omitempty" json:"pool_name,omitempty"`

	// Username for remote media download
	// Required: false
	UsernameDL string `url:"usernameDL,omitempty" json:"usernameDL,omitempty"`

	// Password for remote media download
	// Required: false
	PasswordDl string `url:"passwordDL,omitempty" json:"passwordDL,omitempty"`

	// Binary architecture of this image
	// Should be one of:
	//	- X86_64
	//	- PPC64_LE
	// Required: false
	Architecture string `url:"architecture,omitempty" json:"architecture,omitempty"`

	// List of types of compute suitable for image.
	// Example: [ "KVM_X86" ]
	// Required: true
	Drivers []string `url:"drivers" json:"drivers"`
}

func (irq CreateCDROMImageRequest) validate() error {
	if irq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if irq.URL == "" {
		return errors.New("validation-error: field URL must be set")
	}
	if irq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
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

// CreateCDROMImage creates CD-ROM image from an ISO identified by URL
func (i Image) CreateCDROMImage(ctx context.Context, req CreateCDROMImageRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/image/createCDROMImage"

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
