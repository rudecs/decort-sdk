package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for edit image
type EditRequest struct {
	// ID of the image to edit
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// Name for the image
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Username for the image
	// Required: false
	Username string `url:"username,omitempty" json:"username,omitempty"`

	// Password for the image
	// Required: false
	Password string `url:"password,omitempty" json:"password,omitempty"`

	// Account ID to make the image exclusive
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Does this machine supports hot resize
	// Required: false
	HotResize bool `url:"hotresize,omitempty" json:"hotresize,omitempty"`

	// Does this image boot OS
	// Required: false
	Bootable bool `url:"bootable,omitempty" json:"bootable,omitempty"`
}

func (irq EditRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// Edit edits an existing image
func (i Image) Edit(ctx context.Context, req EditRequest) (bool, error) {
	err := req.validate()
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
