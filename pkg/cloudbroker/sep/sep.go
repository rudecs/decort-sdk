// Operator actions for handling interventions on a storage endpoint provider
package sep

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to storage endpoint provider
type SEP struct {
	client interfaces.Caller
}

// Builder for SEP endpoints
func New(client interfaces.Caller) *SEP {
	return &SEP{
		client: client,
	}
}
