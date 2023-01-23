package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for update resource types in account
type UpdateResourceTypesRequest struct {
	// ID of resource group
	// Required: true
	RGID uint64 `url:"rgId"`

	// Resource types available to create in this resource group
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

func (rgrq UpdateResourceTypesRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if len(rgrq.ResTypes) > 0 {
		for _, value := range rgrq.ResTypes {
			validate := validators.StringInSlice(value, []string{"compute", "vins", "k8s", "openshift", "lb", "flipgroup"})
			if !validate {
				return errors.New("validation-error: Every resource type specified should be one of [compute, vins, k8s, openshift, lb, flipgroup]")
			}
		}
	}

	return nil
}

func (r RG) UpdateResourceTypes(ctx context.Context, req UpdateResourceTypesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/updateResourceTypes"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
