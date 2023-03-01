package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for resize the group
type GroupResizeRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// ID of the Compute Group to resize
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId"`

	// Either delta or absolute value of computes
	// Required: true
	Count int64 `url:"count" json:"count"`

	// Either delta or absolute value of computes
	// Should be one of:
	//	- ABSOLUTE
	//	- RELATIVE
	// Required: true
	Mode string `url:"mode" json:"mode"`
}

func (bsrq GroupResizeRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}
	if bsrq.Mode == "RELATIVE" && bsrq.Count == 0 {
		return errors.New("field Count can not be equal to 0 if Mode if 'RELATIVE'")
	}
	validate := validators.StringInSlice(bsrq.Mode, []string{"RELATIVE", "ABSOLUTE"})
	if !validate {
		return errors.New("field Mode can only be one of 'RELATIVE' or 'ABSOLUTE'")
	}

	return nil
}

// GroupResize resize the group by changing the number of computes
func (b BService) GroupResize(ctx context.Context, req GroupResizeRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/bservice/groupResize"

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
