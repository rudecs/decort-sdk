package k8s

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type CreateRequest struct {
	Name            string   `url:"name"`
	RGID            uint64   `url:"rgId"`
	K8SCIID         uint64   `url:"k8ciId"`
	WorkerGroupName string   `url:"workerGroupName"`
	Labels          []string `url:"labels,omitempty"`
	Taints          []string `url:"taints,omitempty"`
	Annotations     []string `url:"annotations,omitempty"`
	MasterNum       uint     `url:"masterNum,omitempty"`
	MasterCPU       uint     `url:"masterCPU,omitempty"`
	MasterRAM       uint     `url:"masterRam,omitempty"`
	MasterDisk      uint     `url:"masterDisk,omitempty"`
	WorkerNum       uint     `url:"workerNum,omitempty"`
	WorkerCPU       uint     `url:"workerCPU,omitempty"`
	WorkerRAM       uint     `url:"workerRam,omitempty"`
	WorkerDisk      uint     `url:"workerDisk,omitempty"`
	ExtNetID        uint64   `url:"extnetId,omitempty"`
	WithLB          bool     `url:"withLB,omitempty"`
	Description     string   `url:"desc, omitempty"`
}

func (krq CreateRequest) Validate() error {
	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if krq.K8SCIID == 0 {
		return errors.New("validation-error: field K8SCIID can not be empty or equal to 0")
	}

	if krq.WorkerGroupName == "" {
		return errors.New("validation-error: field WorkerGroupName can not be empty")
	}

	return nil
}

func (k8s K8S) Create(ctx context.Context, req CreateRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/k8s/create"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(res), "\"", ""), nil

}
