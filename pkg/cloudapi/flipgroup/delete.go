package flipgroup

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete FLIPGroup
type DeleteRequest struct {
	// FLIPGroup ID
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId" json:"flipgroupId"`
}

func (frq DeleteRequest) validate() error {
	if frq.FLIPGroupID == 0 {
		return errors.New("field FLIPGroupID can not be empty or equal to 0")
	}

	return nil
}

// Delete method wil delete Floating IP group
func (f FLIPGroup) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/delete"

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
