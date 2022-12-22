// API Actors for managing resource groups. These actors are the final API for end users to manage resource groups
package rg

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to resource group
type RG struct {
	client interfaces.Caller
}

// Builder for resource group endpoints
func New(client interfaces.Caller) *RG {
	return &RG{
		client,
	}
}
