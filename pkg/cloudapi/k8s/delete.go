package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	K8SID       uint64 `url:"k8sId"`
	Permanently bool   `url:"permanently"`
}

func (krq DeleteRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/k8s/delete"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}
	return result, nil
}
