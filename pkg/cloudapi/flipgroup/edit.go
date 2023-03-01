package flipgroup

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for edit FLIPGroup
type EditRequest struct {
	// FLIPGroup ID
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId" json:"flipgroupId"`

	// FLIPGroup name
	// Required: true
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// FLIPGroup description
	// Required: true
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

func (frq EditRequest) validate() error {
	if frq.FLIPGroupID == 0 {
		return errors.New("field FLIPGroupID can not be empty or equal to 0")
	}

	return nil
}

// Edit edits FLIPGroup fields
func (f FLIPGroup) Edit(ctx context.Context, req EditRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/edit"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
