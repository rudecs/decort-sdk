package compute

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for insert new CD image
type CDInsertRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of CD-ROM image
	// Required: true
	CDROMID uint64 `url:"cdromId"`

	// Reason to insert
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq CDInsertRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.CDROMID == 0 {
		return errors.New("validation-error: field CDROMID must be set")
	}

	return nil
}

// CDInsert insert new CD image to compute's CD-ROM
func (c Compute) CDInsert(ctx context.Context, req CDInsertRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/cdInsert"

	_, err = c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
