package k8ci

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create K8CI instance
type CreateRequest struct {
	// Name of catalog item
	// Required: true
	Name string `url:"name" json:"name"`

	// Version tag
	// Required: true
	Version string `url:"version" json:"version"`

	// Optional description
	// Required: false
	Description string `url:"description,omitempty" json:"description,omitempty"`

	// Image ID for master K8S node
	// Required: true
	MasterImageID uint64 `url:"masterImageId" json:"masterImageId"`

	// Compute driver
	// Should be one of:
	//	- KVM_X86
	//	- KVM_PPC
	//	- etc
	// Required: true
	MasterDriver string `url:"masterDriver" json:"masterDriver"`

	// Image ID for worker K8S node
	// Required: true
	WorkerImageID uint64 `url:"workerImageId" json:"workerImageId"`

	// Compute driver
	// Should be one of
	//	- KVM_X86
	//	- KVM_PPC
	//	- etc
	// Required: true
	WorkerDriver string `url:"workerDriver" json:"workerDriver"`

	// Image ID for load balancer node
	// Required: true
	LBImageID uint64 `url:"lbImageId" json:"lbImageId"`

	// List of account IDs, which have access to this item.
	// If empty, any account has access
	// Required: false
	SharedWith []uint64 `url:"sharedWith,omitempty" json:"sharedWith,omitempty"`

	// Policy limit on maximum number of master nodes
	// Required: true
	MaxMasterCount uint64 `url:"maxMasterCount" json:"maxMasterCount"`

	// Policy limit on maximum number of worker nodes
	// Required: true
	MaxWorkerCount uint64 `url:"maxWorkerCount" json:"maxWorkerCount"`
}

func (krq CreateRequest) validate() error {
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.Version == "" {
		return errors.New("validation-error: field Version must be set")
	}
	if krq.MasterImageID == 0 {
		return errors.New("validation-error: field MasterImageID must be set")
	}
	if krq.MasterDriver == "" {
		return errors.New("validation-error: field MasterDriver must be set")
	}
	if krq.WorkerImageID == 0 {
		return errors.New("validation-error: field WorkerImageID must be set")
	}
	if krq.WorkerDriver == "" {
		return errors.New("validation-error: field WorkerDriver must be set")
	}
	if krq.LBImageID == 0 {
		return errors.New("validation-error: field LBImageID must be set")
	}
	if krq.MaxMasterCount == 0 {
		return errors.New("validation-error: field MaxMasterCount must be set")
	}
	if krq.MaxWorkerCount == 0 {
		return errors.New("validation-error: field MaxWorkerCount must be set")
	}
	validate := validators.StringInSlice(krq.MasterDriver, []string{"KVM_X86", "KVM_PPC"})
	if !validate {
		return errors.New("validation-error: field MasterDriver must be one of KVM_X86, KVM_PPC")
	}
	validate = validators.StringInSlice(krq.WorkerDriver, []string{"KVM_X86", "KVM_PPC"})
	if !validate {
		return errors.New("validation-error: field WorkerDriver must be one of KVM_X86, KVM_PPC")
	}

	return nil
}

// Create creates a new K8CI instance
func (k K8CI) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/k8ci/create"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
