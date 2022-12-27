package k8s

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for create kubernetes cluster
type CreateRequest struct {
	// Name of Kubernetes cluster
	// Required: true
	Name string `url:"name"`

	// Resource Group ID for cluster placement
	// Required: true
	RGID uint64 `url:"rgId"`

	// ID of Kubernetes catalog item (k8sci) for cluster
	// Required: true
	K8SCIID uint64 `url:"k8ciId"`

	// Name for first worker group created with cluster
	// Required: true
	WorkerGroupName string `url:"workerGroupName"`

	// List of strings with labels for default worker group
	// i.e: ["label1=value1", "label2=value2"]
	// Required: false
	Labels []string `url:"labels,omitempty"`

	// List of strings with taints for default worker group
	// i.e: ["key1=value1:NoSchedule", "key2=value2:NoExecute"]
	// Required: false
	Taints []string `url:"taints,omitempty"`

	// List of strings with annotations for worker group
	// i.e: ["key1=value1", "key2=value2"]
	// Required: false
	Annotations []string `url:"annotations,omitempty"`

	// Number of master nodes to create
	// Required: false
	MasterNum uint `url:"masterNum,omitempty"`

	// Master node CPU count
	// Required: false
	MasterCPU uint `url:"masterCpu,omitempty"`

	// Master node RAM volume in MB
	// Required: false
	MasterRAM uint `url:"masterRam,omitempty"`

	// Master node boot disk size in GB If 0 is specified, size is defined by the OS image size
	// Required: false
	MasterDisk uint `url:"masterDisk,omitempty"`

	// Number of worker nodes to create in default worker group
	// Required: false
	WorkerNum uint `url:"workerNum,omitempty"`

	// Worker node CPU count
	// Required: false
	WorkerCPU uint `url:"workerCpu,omitempty"`

	// Worker node RAM volume in MB
	// Required: false
	WorkerRAM uint `url:"workerRam,omitempty"`

	// Worker node boot disk size in GB. If 0 is specified, size is defined by the OS image size
	// Required: false
	WorkerDisk uint `url:"workerDisk,omitempty"`

	// ID of the external network to connect load balancer and cluster ViNS. If 0 is specified, external network selects automatically to
	// Required: false
	ExtNetID uint64 `url:"extnetId,omitempty"`

	// Create Kubernetes cluster with masters nodes behind load balancer if true.
	// Otherwise give all cluster nodes direct external addresses from selected ExtNet
	// Required: false
	WithLB bool `url:"withLB,omitempty"`

	// Text description of this Kubernetes cluster
	// Required: false
	Description string `url:"desc, omitempty"`
}

func (krq CreateRequest) validate() error {
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

// Create creates a new Kubernetes cluster in the specified Resource Group
func (k8s K8S) Create(ctx context.Context, req CreateRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/k8s/create"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
