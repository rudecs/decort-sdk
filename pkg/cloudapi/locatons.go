package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/locations"

func (ca *CloudApi) Locations() *locations.Locations {
	return locations.New(ca.client)
}
