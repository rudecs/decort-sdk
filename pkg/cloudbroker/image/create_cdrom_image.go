package image

import (
	"context"
	"errors"
	"github.com/rudecs/decort-sdk/internal/validators"
	"net/http"
	"strconv"
)

type CreateCDROMImageRequest struct {
	Name         string   `url:"name"`
	URL          string   `url:"url"`
	GID          uint64   `url:"gid"`
	AccountID    uint64   `url:"accountId,omitempty"`
	SepID        uint64   `url:"sep_id,omitempty"`
	PoolName     string   `url:"pool_name,omitempty"`
	UsernameDL   string   `url:"usernameDL,omitempty"`
	PasswordDl   string   `url:"passwordDL,omitempty"`
	Architecture string   `url:"architecture,omitempty"`
	Drivers      []string `url:"drivers"`
}

func (irq CreateCDROMImageRequest) Validate() error {
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

func (i Image) CreateCDROMImage(ctx context.Context, req CreateCDROMImageRequest) (uint64, error) {
	err := req.Validate()
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
