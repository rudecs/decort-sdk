// API for kubernetes clusters management
package k8s

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to K8S
type K8S struct {
	client interfaces.Caller
}

// Builder for K8S endpoints
func New(client interfaces.Caller) *K8S {
	return &K8S{
		client: client,
	}
}
