package kvmx86

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateBlankRequest struct {
	RGID     uint64 `url:"rgId"`
	Name     string `url:"name"`
	CPU      uint64 `url:"cpu"`
	RAM      uint64 `url:"ram"`
	BootDisk uint64 `url:"bootDisk"`
	SepID    uint64 `url:"sepId"`
	Pool     string `url:"pool"`
	NetType  string `url:"netType,omitempty"`
	NetID    uint64 `url:"netId,omitempty"`
	IPAddr   string `url:"ipAddr,omitempty"`
	Desc     string `url:"desc,omitempty"`
}

func (krq CreateBlankRequest) Validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU must be set")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM must be set")
	}
	if krq.BootDisk == 0 {
		return errors.New("validation-error: field BootDisk must be set")
	}
	if krq.SepID == 0 {
		return errors.New("validation-error: field SepID must be set")
	}
	if krq.Pool == "" {
		return errors.New("validation-error: field Pool must be set")
	}

	return nil
}

func (k KVMX86) CreateBlank(ctx context.Context, req CreateBlankRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/kvmx86/createBlank"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}