package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type UserListRequest struct {
	ComputeId uint64 `url:"computeId "`
}

func (crq UserListRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) UserList(ctx context.Context, req UserListRequest, options ...opts.DecortOpts) (*UserList, error) {

	url := "/compute/userList"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	userList := &UserList{}

	err = json.Unmarshal(res, userList)
	if err != nil {
		return nil, err
	}

	return userList, nil
}
