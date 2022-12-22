package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for insert new CD image
type CDInsertRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of CD-ROM image
	// Required: true
	CDROMID uint64 `url:"cdromId"`
}

func (crq CDInsertRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.CDROMID == 0 {
		return errors.New("validation-error: field CDROMID can not be empty or equal to 0")
	}

	return nil
}

// CDInsert insert new CD image to compute's CD-ROM
func (c Compute) CDInsert(ctx context.Context, req CDInsertRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/cdInsert"

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
