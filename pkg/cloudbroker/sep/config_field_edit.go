package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for edit config fields
type ConfigFieldEditRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// Field name
	// Required: true
	FieldName string `url:"field_name"`

	// Field value
	// Required: true
	FieldValue string `url:"field_value"`

	// Field type
	// Should be one of:
	//	- int
	//	- str
	//	- bool
	//	- list
	//	- dict
	// Required: true
	FieldType string `url:"field_type"`
}

func (srq ConfigFieldEditRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.FieldName == "" {
		return errors.New("validation-error: field FieldName must be set")
	}
	if srq.FieldValue == "" {
		return errors.New("validation-error: field FieldValue must be set")
	}
	if srq.FieldType == "" {
		return errors.New("validation-error: field FieldType must be set")
	}
	validate := validators.StringInSlice(srq.FieldType, []string{"int", "str", "bool", "list", "dict"})
	if !validate {
		return errors.New("validation-error: field FieldType must be one of int, str, bool, list, dict")
	}

	return nil
}

// ConfigFieldEdit edit SEP config field value
func (s SEP) ConfigFieldEdit(ctx context.Context, req ConfigFieldEditRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/configFieldEdit"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
