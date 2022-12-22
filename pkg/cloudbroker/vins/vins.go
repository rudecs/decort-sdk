// API Actor for managing VINS. This actor is a final API for endusers to manage VINS
package vins

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to VINS
type VINS struct {
	client interfaces.Caller
}

// Builder for VINS endpoints
func New(client interfaces.Caller) *VINS {
	return &VINS{
		client: client,
	}
}
