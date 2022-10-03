package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/rg"

func (ca *CloudApi) RG() *rg.RG {
	return rg.New(ca.client)
}
