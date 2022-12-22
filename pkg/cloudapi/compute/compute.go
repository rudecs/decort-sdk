// API Actor for managing Compute. This actor is a final API for endusers to manage Compute
package compute

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to compute
type Compute struct {
	client interfaces.Caller
}

// Builder for compute endpoints
func New(client interfaces.Caller) *Compute {
	return &Compute{
		client,
	}
}
