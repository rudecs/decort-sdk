package decortsdk

import "github.com/rudecs/decort-sdk/pkg/cloudapi/locations"

func (dc *Client) Locations() *locations.Locations {
	return locations.New(dc)
}
