package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/locations"

// Accessing the Locations method group
func (ca *CloudAPI) Locations() *locations.Locations {
	return locations.New(ca.client)
}
