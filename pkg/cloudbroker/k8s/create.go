package k8s

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for create K8S
type CreateRequest struct {
	// Name of kubernetes cluster
	// Required: true
	Name string `url:"name" json:"name"`

	// Resource group ID for cluster placement
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// ID of kubernetes catalog item (K8CI) for cluster
	// Required: true
	K8CIID uint64 `url:"k8ciId" json:"k8ciId"`

	// Name for first worker group created with cluster
	// Required: true
	WorkerGroupName string `url:"workerGroupName" json:"workerGroupName"`

	// ID of SEP to create boot disks for master nodes.
	// Uses images SEP ID if not set
	// Required: false
	MasterSEPID uint64 `url:"masterSepId,omitempty" json:"masterSepId,omitempty"`

	// Pool to use if master SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	MasterSEPPool string `url:"masterSepPool,omitempty" json:"masterSepPool,omitempty"`

	// ID of SEP to create boot disks for default worker nodes group.
	// Uses images SEP ID if not set
	// Required: false
	WorkerSEPID uint64 `url:"workerSepId,omitempty" json:"workerSepId,omitempty"`

	// Pool to use if worker SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	WorkerSEPPool string `url:"workerSepPool,omitempty" json:"workerSepPool,omitempty"`

	// List of strings with labels for default worker group
	// i.e: ["label1=value1", "label2=value2"]
	// Required: false
	Labels []string `url:"labels,omitempty" json:"labels,omitempty"`

	// List of strings with taints for default worker group
	// i.e: ["key1=value1:NoSchedule", "key2=value2:NoExecute"]
	// Required: false
	Taints []string `url:"taints,omitempty" json:"taints,omitempty"`

	// List of strings with annotations for worker group
	// i.e: ["key1=value1", "key2=value2"]
	// Required: false
	Annotations []string `url:"annotations,omitempty" json:"annotations,omitempty"`

	// Number of master nodes to create
	// Required: false
	MasterNum uint64 `url:"masterNum,omitempty" json:"masterNum,omitempty"`

	// Master node CPU count
	// Required: false
	MasterCPU uint64 `url:"masterCpu,omitempty" json:"masterCpu,omitempty"`

	// Master node RAM volume in MB
	// Required: false
	MasterRAM uint64 `url:"masterRam,omitempty" json:"masterRam,omitempty"`

	// Master node boot disk size in GB If 0 is specified, size is defined by the OS image size
	// Required: false
	MasterDisk uint64 `url:"masterDisk,omitempty" json:"masterDisk,omitempty"`

	// Number of worker nodes to create in default worker group
	// Required: false
	WorkerNum uint64 `url:"workerNum,omitempty" json:"workerNum,omitempty"`

	// Worker node CPU count
	// Required: false
	WorkerCPU uint64 `url:"workerCpu,omitempty" json:"workerCpu,omitempty"`

	// Worker node RAM volume in MB
	// Required: false
	WorkerRAM uint64 `url:"workerRam,omitempty" json:"workerRam,omitempty"`

	// Worker node boot disk size in GB. If 0 is specified, size is defined by the OS image size
	// Required: false
	WorkerDisk uint64 `url:"workerDisk,omitempty" json:"workerDisk,omitempty"`

	// ID of the external network to connect load balancer and cluster VINS. If 0 is specified, external network selects automatically to
	// Required: false
	ExtNetID uint64 `url:"extnetId,omitempty" json:"extnetId,omitempty"`

	// Create kubernetes cluster with masters nodes behind load balancer if true.
	// Otherwise give all cluster nodes direct external addresses from selected external network
	// Required: false
	WithLB bool `url:"withLB,omitempty" json:"withLB,omitempty"`

	// Text description of this kubernetes cluster
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

func (krq CreateRequest) validate() error {
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if krq.K8CIID == 0 {
		return errors.New("validation-error: field K8CIID must be set")
	}
	if krq.WorkerGroupName == "" {
		return errors.New("validation-error: field WorkerGroupName must be set")
	}

	return nil
}

// Create creates a new kubernetes cluster in the specified resource group
func (k K8S) Create(ctx context.Context, req CreateRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/k8s/create"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
