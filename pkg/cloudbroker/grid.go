package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudbroker/grid"

// Accessing the Grid method group
func (cb *CloudBroker) Grid() *grid.Grid {
	return grid.New(cb.client)
}
