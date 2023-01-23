package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for update resource types in account
type UpdateResourceTypesRequest struct {
	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Resource types available to create in this account
	// Each element in a resource type slice must be one of:
	//	- compute
	//	- vins
	//	- k8s
	//	- openshift
	//	- lb
	//	- flipgroup
	// Required: true
	ResTypes []string `url:"resourceTypes"`
}

func (arq UpdateResourceTypesRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if len(arq.ResTypes) > 0 {
		for _, value := range arq.ResTypes {
			validate := validators.StringInSlice(value, []string{"compute", "vins", "k8s", "openshift", "lb", "flipgroup"})
			if !validate {
				return errors.New("validation-error: Every resource type specified should be one of [compute, vins, k8s, openshift, lb, flipgroup]")
			}
		}
	}

	return nil
}

func (a Account) UpdateResourceTypes(ctx context.Context, req UpdateResourceTypesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/updateResourceTypes"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
