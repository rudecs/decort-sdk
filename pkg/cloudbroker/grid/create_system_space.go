package grid

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for create system space
type CreateSystemSpaceRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"id" json:"id"`

	// Name of the account/cloudspace to be created for the system
	// Required: true
	Name string `url:"name" json:"name"`

	// ID of the specific image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// Size of base volume
	// Required: true
	BootSize uint64 `url:"bootsize" json:"bootsize"`

	// Data disk size in gigabytes
	// Required: true
	DataDiskSize uint64 `url:"dataDiskSize" json:"dataDiskSize"`

	// ID of the specific size
	// Required: false
	SizeID uint64 `url:"sizeId,omitempty" json:"sizeId,omitempty"`

	// Number of vcpus to provide
	// Required: false
	VCPUS uint64 `url:"vcpus,omitempty" json:"vcpus,omitempty"`

	// Amount of memory to provide
	// Required: false
	Memory uint64 `url:"memory,omitempty" json:"memory,omitempty"`
}

func (grq CreateSystemSpaceRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if grq.BootSize == 0 {
		return errors.New("validation-error: field BootSize must be set")
	}
	if grq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if grq.DataDiskSize == 0 {
		return errors.New("validation-error: field DataDiskSize must be set")
	}

	return nil
}

// CreateSystemSpace creates system space
func (g Grid) CreateSystemSpace(ctx context.Context, req CreateSystemSpaceRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/grid/createSystemSpace"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
