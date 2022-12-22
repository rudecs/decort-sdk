package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for restores BasicService instance
type RestoreRequest struct {
	// ID of the BasicService to be restored
	// Required: true
	ServiceID uint64 `url:"serviceId"`
}

func (bsrq RestoreRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Restore restores BasicService instance
func (b BService) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/restore"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
