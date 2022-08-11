package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateRequest struct {
	AccountId   uint64 `url:"accountId"`
	GID         uint64 `url:"gid"`
	Name        string `url:"name"`
	Description string `url:"description,omitempty"`
	Size        uint64 `url:"size,omitempty"`
	Type        string `url:"type"`
	SSDSize     uint64 `url:"ssdSize,omitempty"`
	IOps        uint64 `url:"iops"`
	SepId       uint64 `url:"sep_id,omitempty"`
	Pool        string `url:"pool,omitempty"`
}

func (drq CreateRequest) Validate() error {
	if drq.AccountId == 0 {
		return errors.New("validation-error: field AccountId can not be empty or equal to 0")
	}
	if drq.GID == 0 {
		return errors.New("validation-error: field Gid can not be empty or equal to 0")
	}
	if drq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	validType := validators.StringInSlice(drq.Type, []string{"B", "D", "T"})
	if !validType {
		return errors.New("validation-error: field Type must be set as B, D or T")
	}

	if drq.IOps == 0 {
		return errors.New("validation-error: field IOps must be set")
	}

	return nil
}

func (d Disks) Create(ctx context.Context, req CreateRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/disks/create"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := d.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil

}
