package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type UserGrantRequest struct {
	ComputeID  uint64 `url:"computeId"`
	Username   string `url:"userName"`
	AccessType string `url:"accesstype"`
}

func (crq UserGrantRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.Username == "" {
		return errors.New("validation-error: field UserName can not be empty")
	}

	if crq.AccessType == "" {
		return errors.New("validation-error: field AccessType can not be empty")
	}
	validator := validators.StringInSlice(crq.AccessType, []string{"R", "RCX", "ARCXDU"})
	if !validator {
		return errors.New("validation-error: field AccessType can be only R, RCX or ARCXDU")
	}

	return nil
}

func (c Compute) UserGrant(ctx context.Context, req UserGrantRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/userGrant"

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
