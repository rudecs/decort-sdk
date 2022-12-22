package k8s

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for add workers group
type WorkersGroupAddRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// Worker group name
	// Required: true
	Name string `url:"name"`

	// ID of SEP to create boot disks for default worker nodes group.
	// Uses images SEP ID if not set
	// Required: false
	WorkerSEPID uint64 `url:"workerSepId,omitempty"`

	// Pool to use if worker SEP ID is set, can be also empty if
	// needed to be chosen by system
	// Required: false
	WorkerSEPPool string `url:"workerSepPool,omitempty"`

	// List of strings with labels for worker group
	// i.e: ["label1=value1", "label2=value2"]
	// Required: false
	Labels []string `url:"labels,omitempty"`

	// List of strings with taints for worker group
	// i.e: ["key1=value1:NoSchedule", "key2=value2:NoExecute"]
	// Required: false
	Taints []string `url:"taints,omitempty"`

	// List of strings with annotations for worker group
	// i.e: ["key1=value1", "key2=value2"]
	// Required: false
	Annotations []string `url:"annotations,omitempty"`

	// Number of worker nodes to create
	// Required: false
	WorkerNum uint64 `url:"workerNum,omitempty"`

	// Worker node CPU count
	// Required: false
	WorkerCPU uint64 `url:"workerCpu,omitempty"`

	// Worker node RAM volume in MB
	// Required: false
	WorkerRAM uint64 `url:"workerRam,omitempty"`

	// Worker node boot disk size in GB If 0 is specified, size is defined by the OS image size
	// Required: false
	WorkerDisk uint64 `url:"workerDisk,omitempty"`
}

func (krq WorkersGroupAddRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// WorkersGroupAdd adds workers group to kubernetes cluster
func (k K8S) WorkersGroupAdd(ctx context.Context, req WorkersGroupAddRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/k8s/workersGroupAdd"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
