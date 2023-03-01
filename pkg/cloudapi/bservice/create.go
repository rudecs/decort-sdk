package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for BasicService
type CreateRequest struct {
	// Name of the service
	// Required: true
	Name string `url:"name" json:"name"`

	// ID of the Resource Group where this service will be placed
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// Name of the user to deploy SSH key for. Pass empty string if no SSH key deployment is required
	// Required: false
	SSHUser string `url:"sshUser,omitempty" json:"sshUser,omitempty"`

	// SSH key to deploy for the specified user. Same key will be deployed to all computes of the service
	// Required: false
	SSHKey string `url:"sshKey,omitempty" json:"sshKey,omitempty"`
}

func (bsrq CreateRequest) validate() error {
	if bsrq.Name == "" {
		return errors.New("field Name can not be empty")
	}
	if bsrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

// Create creates blank BasicService instance
func (b BService) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/bservice/create"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
