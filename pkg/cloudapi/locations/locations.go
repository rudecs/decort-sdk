// API Actor api for managing locations
package locations

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to locations
type Locations struct {
	client interfaces.Caller
}

// Builder for locations endpoints
func New(client interfaces.Caller) *Locations {
	return &Locations{
		client,
	}
}
