package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type UserListRequest struct {
	ComputeID uint64 `url:"computeId "`
}

func (crq UserListRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) UserList(ctx context.Context, req UserListRequest) (*UserList, error) {

	url := "/cloudapi/compute/userList"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
