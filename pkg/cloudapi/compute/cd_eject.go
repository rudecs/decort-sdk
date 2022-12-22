package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for eject CD image
type CDEjectRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq CDEjectRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// CDEject eject CD image to compute's CD-ROM
func (c Compute) CDEject(ctx context.Context, req CDEjectRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/cdEject"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}
	return result, nil
}
