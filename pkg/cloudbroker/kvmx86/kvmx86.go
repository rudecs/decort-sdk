// API to manage KVM x86 compute instances (x86 VMs)
package kvmx86

import "github.com/rudecs/decort-sdk/interfaces"

// Structure for creating request to KVMX86
type KVMX86 struct {
	client interfaces.Caller
}

// Builder for KVMX86 endpoints
func New(client interfaces.Caller) *KVMX86 {
	return &KVMX86{
		client: client,
	}
}
