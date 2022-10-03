package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type FindGroupByLabelRequest struct {
	K8SID  uint64   `url:"k8sId"`
	Labels []string `url:"labels"`
	Strict bool     `url:"strict"`
}

func (krq FindGroupByLabelRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}

	if len(krq.Labels) == 0 {
		return errors.New("validation-error: field Labels can not be empty")
	}

	return nil
}

func (k8s K8S) FindGroupByLabel(ctx context.Context, req FindGroupByLabelRequest) (K8SGroupList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/k8s/findGroupByLabel"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	k8sGroupList := K8SGroupList{}

	err = json.Unmarshal(res, &k8sGroupList)
	if err != nil {
		return nil, err
	}

	return k8sGroupList, nil
}
