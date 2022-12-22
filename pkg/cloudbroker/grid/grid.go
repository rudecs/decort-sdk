// Operator actions for handling interventions on a grid
package grid

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to grid
type Grid struct {
	client interfaces.Caller
}

// Builder for grid endpoints
func New(client interfaces.Caller) *Grid {
	return &Grid{
		client: client,
	}
}
