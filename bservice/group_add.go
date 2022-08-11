package bservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GroupAddRequest struct {
	ServiceID    uint64   `url:"serviceId"`
	Name         string   `url:"name"`
	Count        uint64   `url:"count"`
	CPU          uint64   `url:"cpu"`
	RAM          uint64   `url:"ram"`
	Disk         uint64   `url:"disk"`
	ImageID      uint64   `url:"imageId"`
	Driver       string   `url:"driver"`
	Role         string   `url:"role,omitempty"`
	VINSes       []uint64 `url:"vinses,omitempty"`
	Extnets      []uint64 `url:"extnets,omitempty"`
	TimeoutStart uint64   `url:"timeoutStart"`
}

func (bsrq GroupAddRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.Name == "" {
		return errors.New("field Name can not be empty")
	}

	if bsrq.Count == 0 {
		return errors.New("field Count can not be empty or equal to 0")
	}

	if bsrq.CPU == 0 {
		return errors.New("field CPU can not be empty or equal to 0")
	}

	if bsrq.RAM == 0 {
		return errors.New("field RAM can not be empty or equal to 0")
	}

	if bsrq.Disk == 0 {
		return errors.New("field Disk can not be empty or equal to 0")
	}

	if bsrq.ImageID == 0 {
		return errors.New("field ImageID can not be empty or equal to 0")
	}

	if bsrq.Driver == "" {
		return errors.New("field Driver can not be empty")
	}

	return nil
}

func (b BService) GroupAdd(ctx context.Context, req GroupAddRequest, options ...opts.DecortOpts) (uint64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	url := "/cloudapi/bservice/groupAdd"
	res, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(string(res), 10, 64)
}
