package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create disk
type CreateRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// ID of the grid (platform)
	// Required: true
	GID uint64 `url:"gid"`

	// Name of disk
	// Required: true
	Name string `url:"name"`

	// Description of disk
	// Required: false
	Description string `url:"description,omitempty"`

	// Size in GB, default is 0
	// Required: false
	Size uint64 `url:"size,omitempty"`

	// Type of disk
	//	- B=Boot
	//	- D=Data
	//	- T=Temp
	// Required: true
	Type string `url:"type"`

	// Size in GB default is 0
	// Required: false
	SSDSize uint64 `url:"ssdSize,omitempty"`

	// Max IOPS disk can perform defaults to 2000
	// Required: false
	IOPS uint64 `url:"iops,omitempty"`

	// Storage endpoint provider ID to create disk
	// Required: false
	SEPID uint64 `url:"sep_id,omitempty"`

	// Pool name to create disk
	// Required: false
	Pool string `url:"pool,omitempty"`
}

func (drq CreateRequest) validate() error {
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

// Create creates a disk
func (d Disks) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
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
