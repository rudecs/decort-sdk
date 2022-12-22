// API Actor for use external networks
package extnet

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

// Structure for creating request to extnet
type ExtNet struct {
	client interfaces.Caller
}

// Builder for extnet endpoints
func New(client interfaces.Caller) *ExtNet {
	return &ExtNet{
		client,
	}
}
