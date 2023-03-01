package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update nodes
type UpdateNodesRequest struct {
	// Image ID
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// List of stacks
	// Required: false
	EnabledStacks []uint64 `url:"enabledStacks,omitempty" json:"enabledStacks,omitempty"`
}

func (irq UpdateNodesRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// UpdateNodes udates image availability on nodes
func (i Image) UpdateNodes(ctx context.Context, req UpdateNodesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/updateNodes"

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
