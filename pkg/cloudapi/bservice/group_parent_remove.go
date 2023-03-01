package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for remove parent Compute Group
// relation to the specified Compute Group
type GroupParentRemoveRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// ID of the Compute Group
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId"`

	// ID of the parent Compute Group
	// to remove from the current Compute Group
	// Required: true
	ParentID uint64 `url:"parentId" json:"parentId"`
}

func (bsrq GroupParentRemoveRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}
	if bsrq.ParentID == 0 {
		return errors.New("field ParentID can not be empty or equal to 0")
	}

	return nil
}

// GroupParentRemove removes parent Compute Group relation to the specified Compute Group
func (b BService) GroupParentRemove(ctx context.Context, req GroupParentRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupParentRemove"

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
