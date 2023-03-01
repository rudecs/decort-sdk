package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete basic service
type DeleteRequest struct {
	// ID of the BasicService to be delete
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// If set to False, Basic service will be deleted to recycle bin. Otherwise destroyed immediately
	// Required: true
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

func (bsrq DeleteRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Delete deletes BasicService instance
func (b BService) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/delete"

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
