package client

import "github.com/rudecs/decort-sdk/locations"

func (dc *decortClient) Locations() *locations.Locations {
	return locations.New(dc)
}
