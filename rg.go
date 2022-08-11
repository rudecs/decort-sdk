package decortsdk

import "github.com/rudecs/decort-sdk/pkg/cloudapi/rg"

func (dc *decortClient) RG() *rg.RG {
	return rg.New(dc)
}
