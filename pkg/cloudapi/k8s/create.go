package k8s

import (
	"context"
	"errors"
	"strings"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateRequest struct {
	Name            string   `url:"name"`
	RGID            uint64   `url:"rgId"`
	K8SCIId         uint64   `url:"k8ciId"`
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
	ExtnetId        uint64   `url:"extnetId,omitempty"`
	WithLB          bool     `url:"withLB,omitempty"`
	Description     string   `url:"desc, omitempty"`
}

func (krq CreateRequest) Validate() error {
	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if krq.RGID == 0 {
		return errors.New("validation-error: field RgId can not be empty or equal to 0")
	}
	if krq.K8SCIId == 0 {
		return errors.New("validation-error: field K8SCIId can not be empty or equal to 0")
	}

	if krq.WorkerGroupName == "" {
		return errors.New("validation-error: field WorkerGroupName can not be empty")
	}

	return nil
}

func (k8s K8S) Create(ctx context.Context, req CreateRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/k8s/create"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := k8s.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(res), "\"", ""), nil

}
