// API to manage KVM PowerPC compute instances (PPC VMs)
package kvmppc

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to KVMPPC
type KVMPPC struct {
	client interfaces.Caller
}

// Builder for KVMPPC endpoints
func New(client interfaces.Caller) *KVMPPC {
	return &KVMPPC{
		client,
	}
}
