package kvmx86

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type MassCreateRequest struct {
	RGID     uint64 `url:"rgId"`
	Name     string `url:"name"`
	Count    uint64 `url:"count"`
	CPU      uint64 `url:"cpu"`
	RAM      uint64 `url:"ram"`
	ImageID  uint64 `url:"imageId"`
	BootDisk uint64 `url:"bootDisk,omitempty"`
	SepID    uint64 `url:"sepId,omitempty"`
	Pool     string `url:"pool,omitempty"`
	NetType  string `url:"netType,omitempty"`
	NetID    uint64 `url:"netId,omitempty"`
	IPAddr   string `url:"ipAddr,omitempty"`
	UserData string `url:"userdata,omitempty"`
	Desc     string `url:"desc,omitempty"`
	Start    bool   `url:"start,omitempty"`
	Reason   string `url:"reason,omitempty"`
}

func (krq MassCreateRequest) Validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.Count == 0 {
		return errors.New("validation-error: field Count must be set")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU must be set")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM must be set")
	}
	if krq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (k KVMX86) MassCreate(ctx context.Context, req MassCreateRequest) ([]uint64, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/kvmx86/massCreate"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	computes := make([]uint64, 0)

	err = json.Unmarshal(res, &computes)
	if err != nil {
		return nil, err
	}

	return computes, nil
}
