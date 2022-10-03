package disks

import (
	"context"
	"errors"
	"github.com/rudecs/decort-sdk/internal/validators"
	"net/http"
	"strconv"
)

type CreateRequest struct {
	AccountID   uint64 `url:"accountId"`
	GID         uint64 `url:"gid"`
	Name        string `url:"name"`
	Description string `url:"description,omitempty"`
	Size        uint64 `url:"size,omitempty"`
	Type        string `url:"type"`
	SSDSize     uint64 `url:"ssdSize"`
	IOPS        uint64 `url:"iops,omitempty"`
	SepID       uint64 `url:"sepId,omitempty"`
	Pool        string `url:"pool,omitempty"`
}

func (drq CreateRequest) Validate() error {
	if drq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	if drq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}

	if drq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	validate := validators.StringInSlice(drq.Type, []string{"B", "D", "T"})
	if !validate {
		return errors.New("validation-error: field must be B, D or T")
	}

	return nil
}

func (d Disks) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/disks/create"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}