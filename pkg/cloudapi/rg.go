package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/rg"

// Accessing the RG method group
func (ca *CloudAPI) RG() *rg.RG {
	return rg.New(ca.client)
}
