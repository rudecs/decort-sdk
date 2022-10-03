package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudapi/locations"

func (cb *CloudBroker) Locations() *locations.Locations {
	return locations.New(cb.client)
}
