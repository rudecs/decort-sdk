package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type EditRequest struct {
	ImageID   uint64 `url:"imageId"`
	Name      string `url:"name,omitempty"`
	Username  string `url:"username,omitempty"`
	Password  string `url:"password,omitempty"`
	AccountID uint64 `url:"accountId,omitempty"`
	HotResize bool   `url:"hotresize,omitempty"`
	Bootable  bool   `url:"bootable,omitempty"`
}

func (irq EditRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (i Image) Edit(ctx context.Context, req EditRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/edit"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
